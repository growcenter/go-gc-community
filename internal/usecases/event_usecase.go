package usecases

import (
	"go-gc-community/internal/models"
	"go-gc-community/internal/repositories"
	"strconv"
	"time"
)

type Event interface {
	GetAll() ([]*models.Events, error)
	GetSessionByEvent(id int) ([]*models.Sessions, time.Time, error)
}

type eventUsecase struct {
	er repositories.Event
	sr repositories.Session
}

func NewEventUsecase(er repositories.Event, sr repositories.Session) *eventUsecase {
	return &eventUsecase{er: er, sr: sr}
}

func (eu *eventUsecase) GetAll() ([]*models.Events, time.Time, error) {
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

func (eu *eventUsecase) GetSessionByEvent(id string) ([]*models.Sessions, *models.Events, time.Time, error) {
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