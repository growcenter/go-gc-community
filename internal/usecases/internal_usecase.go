package usecases

import (
	"errors"
	"fmt"
	"go-gc-community/internal/models"
	"go-gc-community/internal/repositories"
	"strings"
)

type Internal interface {
	SetRole(request *models.SetRoleRequest) (*models.User, error)
	List(page, limit, sort, filter, accountNumber string) ([]*models.Registrations, error)
	Verify(code, accountNumber string) (*models.Registrations, *models.Events, *models.Registrations, error)
}

type internalUsecase struct {
	ur repositories.User
	er repositories.Event
	sr repositories.Session
	rr repositories.Registration
}

func NewInternalUsecase(ur repositories.User, er repositories.Event, sr repositories.Session, rr repositories.Registration) *internalUsecase {
	return &internalUsecase{ur: ur, er: er, sr: sr, rr: rr}
}

func (iu *internalUsecase) SetRole(request *models.SetRoleRequest) (*models.User, error) {
	current, err := iu.ur.Find("email", strings.ToLower(request.Email))
	if err != nil {
		return nil, err
	}

	if current.ID == 0 {
		return nil, errors.New("user does not exist")
	}

	if current.AccountNumber != request.AccountNumber {
		return nil, errors.New("account number is not valid")
	}

	if request.RoleId != "02" && request.RoleId != "01" {
		return nil, errors.New("currently other role does not exist")
	}

	if request.RoleId == current.RoleId {
		return nil, fmt.Errorf("No changes, your current role is %s", current.RoleId)
	}

	current.RoleId = request.RoleId
	change, err := iu.ur.Update(current)
	if err != nil {
		return nil, err
	}

	return change, nil
}

func (iu *internalUsecase) List(page, limit, sort, filter, accountNumber string) ([]*models.Registrations, error) {
	user, err := iu.ur.Find("account_number", accountNumber)
	if err != nil {
		return nil, err
	}

	if user.RoleId != "02" {
		return nil, errors.New("role is not allowed")
	}

	result, err := iu.rr.List(page, limit, sort, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (iu *internalUsecase) Verify(code, accountNumber string) (*models.Registrations, *models.Events, *models.Sessions, error) {
	user, err := iu.ur.Find("account_number", accountNumber)
	if err != nil {
		return nil, nil, nil, err
	}

	if user.RoleId != "02" {
		return nil, nil, nil, errors.New("role is not allowed")
	}

	current, err := iu.rr.Find("code", code)
	if err != nil {
		return nil, nil, nil, err
	}

	if current.ID == 0 {
		return nil, nil, nil, errors.New("registration not found")
	}

	if current.Status != "02" && current.Status != "01" && current.Status != "00" {
		return nil, nil, nil, errors.New("your code is invalid")
	}

	if current.Status == "02" {
		return nil, nil, nil, errors.New("your registration is already cancelled previously")
	}

	if current.Status == "00" {
		return nil, nil, nil, errors.New("your code is already verified")
	}

	event, err := iu.er.Find("id", current.EventsId)
	if err != nil {
		return nil, nil, nil, err
	}

	session, err := iu.sr.Find("id", current.SessionsId)
	if err != nil {
		return nil, nil, nil, err
	}

	current.Status = "00"
	current.BookedBy = strings.ToLower(user.Email)
	session.ScannedSeats += 1
	session.UnscannedSeats -= 1

	updatedReg, err := iu.rr.Update(current)
	if err != nil {
		return nil, nil, nil, err
	}

	updatedSes, err := iu.sr.Update(session)
	if err != nil {
		return nil, nil, nil, err
	}	


	return updatedReg, event, updatedSes, nil
}