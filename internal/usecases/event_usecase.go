package usecases

import (
	"errors"
	"fmt"
	"go-gc-community/internal/models"
	"go-gc-community/internal/repositories"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Event interface {
	Events() ([]*models.Events, error)
	Sessions(id int) ([]*models.Sessions, time.Time, error)
	Register(request *models.RegistrationRequest) (*models.Registrations, []*models.Registrations, bool, error)
}

type eventUsecase struct {
	ur repositories.User
	er repositories.Event
	sr repositories.Session
	rr repositories.Registration
}

func NewEventUsecase(ur repositories.User, er repositories.Event, sr repositories.Session, rr repositories.Registration) *eventUsecase {
	return &eventUsecase{ur: ur, er: er, sr: sr, rr: rr}
}

func (eu *eventUsecase) Event(id int) (*models.Events, error) {
	event, err := eu.er.Find("id", id)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (eu *eventUsecase) Events() ([]*models.Events, time.Time, error) {
	currentTime := time.Now()
	
	openChanges := eu.er.UpdateFilter("open_registration < ?", currentTime, "status", "OPEN")
	err := openChanges.Error
	if err != nil {
		return nil, time.Now(), err
	}

	closeChanges := eu.er.UpdateFilter("closed_registration < ?", currentTime, "status", "CLOSED")
	err = closeChanges.Error
	if err != nil {
		return nil, time.Now(), err
	}


	event, err := eu.er.All()
	if err != nil {
		return nil, time.Now(), err
	}

	return event, currentTime, nil
}

func (eu *eventUsecase) Sessions(id string) ([]*models.Sessions, *models.Events, time.Time, error) {
	var e *models.Events
	currentTime := time.Now()
	eventId, err := strconv.Atoi(id)
	if err != nil {
		return nil, e, currentTime, err
	}

	openChanges := eu.sr.UpdateFilter("open_registration < ?", currentTime, "status", "OPEN")
	err = openChanges.Error
	if err != nil {
		return nil, e, currentTime, err
	}

	closeChanges := eu.sr.UpdateFilter("closed_registration < ?", currentTime, "status", "CLOSED")
	err = closeChanges.Error
	if err != nil {
		return nil, e, currentTime, err
	}

	event, err := eu.er.Find("id", eventId)
	if err != nil {
		return nil, event, currentTime, err
	}

	session, err := eu.sr.AllWithFilter("events_id", eventId)
	if err != nil {
		return nil, event, currentTime, err
	}

	return session, event, currentTime, err
}

func (eu *eventUsecase) Register(request *models.RegistrationRequest) (*models.Registrations, []*models.Registrations, bool, int, error) {
	// Retrieve Event Data by EventId
	event, err := eu.er.Find("id", request.EventID)
	if err != nil {
		return nil, nil, false, 0, err
	}

	// Validation if Event Status is already closed
	if event.Status == "CLOSED" {
		return nil, nil, false, 0, errors.New("cannot registed as the event is already closed")
	}
	
	// Validation if the current time is not the time yet for the event to open registration
	if event.OpenRegistration.After(time.Now()) {
		return nil, nil, false, 0, errors.New("event registration period is not open yet")
	}

	// Validation if the current time is already over the close register time
	if event.ClosedRegistration.Before(time.Now()) {
		return nil, nil, false, 0, errors.New("event registration period is already closed")
	}

	// Retrieve Session Data by SessionId
	session, err := eu.sr.Find("id", request.SessionID)
	if err != nil {
		return nil, nil, false, 0, err
	}

	// Validation if Session Status is already closed
	if session.Status == "CLOSED" {
		return nil, nil, false, 0, errors.New("cannot registed as the session is already closed")
	}

	// Validation if the current time is not the time yet for the session to open registration
	if session.OpenRegistration.After(time.Now()) {
		return nil, nil, false, 0, errors.New("session registration period is not open yet")
	}

	// Validation if the current time is already over the close register time
	if session.ClosedRegistration.Before(time.Now()) {
		return nil, nil, false, 0, errors.New("session registration period is already closed")
	}

	// Count how many user register from the request (main user + count of array)
	count := 1 + len(request.Others)

	// Validation if the number of requested seat is more than allowable seating
	if count > session.MaxSeating {
		return nil, nil, false, 0, fmt.Errorf("you cannot enter more than %d user", session.MaxSeating)
	}

	// Validation if the session seating is unavilable anymore
	if session.AvailableSeats == 0 {
		return nil, nil, false, 0, errors.New("no seats left on this session")
	}

	isRegistered, err := eu.rr.Find("email", strings.ToLower(request.MainEmail))
	if err != nil {
		return nil, nil, false, 0, err
	}

	isBooked, err := eu.rr.Find("booked_by", strings.ToLower(request.MainEmail))
	if err != nil {
		return nil, nil, false, 0, err
	}

	isAccount, err := eu.ur.Find("email", strings.ToLower(request.MainEmail))
	if err != nil {
		return nil, nil, false, 0, errors.New("user does not have account yet")
	}

	// Check if user already registered or not
	if isRegistered.ID != 0 || strings.ToLower(isRegistered.Email) == strings.ToLower(request.MainEmail) {
		return nil, nil, false, 0, errors.New("user already registered")
	}
	
	// Check if the user already booked for other, means user can only register once
	if isBooked.ID != 0 || strings.ToLower(isBooked.Email) == strings.ToLower(request.MainEmail) {
		return nil, nil, false, 0, errors.New(fmt.Sprintf("You are already registered by: %s", strings.ToUpper(request.MainName)))
	}

	if isAccount.ID == 0 || strings.ToLower(isAccount.Email) != strings.ToLower(request.MainEmail) {
		return nil, nil, false, 0, errors.New("user does not have account yet")
	}

	Reg := models.Registrations{
		Name: strings.ToUpper(request.MainName),
		Email: strings.ToLower(isAccount.Email),
		Code: (uuid.New()).String(),
		EventsId: event.ID,
		SessionsId: session.ID,
		Status: "01",
		BookedBy: strings.ToLower(isAccount.Email),
		AccountNumber: isAccount.AccountNumber,
	}

	main, err := eu.rr.Create(&Reg)
	if err != nil {
		return nil, nil, false, 0, err
	}

	var others []*models.Registrations
	for _, other := range request.Others {
		isRegistered, err := eu.rr.Find("email", strings.ToLower(other.Email))
		if err != nil {
			return nil, nil, false, 0, err
		}
		if isRegistered.ID != 0 || strings.ToLower(isRegistered.Email) == strings.ToLower(other.Email) {
			return nil, nil, false, 0, errors.New("user already registered")
		}

		otherReg := models.Registrations {
			Name: strings.ToUpper(other.Name),
			Email: strings.ToLower(other.Email),
			Code: (uuid.New()).String(),
			EventsId: request.EventID,
			SessionsId: request.SessionID,
			Status: "01",
			BookedBy: strings.ToLower(request.MainEmail),
		}

		secondary, err := eu.rr.Create(&otherReg)
		if err != nil {
			return nil, nil, false, 0, err
		}

		others = append(others, secondary)
	}

	session.AvailableSeats -= count
	session.BookedSeats += count
	session.UnscannedSeats += count
	session.TotalRegistration += count

	_, err = eu.sr.Update(session)
	if err != nil {
		return nil, nil, false, 0, err
	}

	return main, others, true, count, nil
}