package usecases

import (
	"go-gc-community/internal/models"
	"go-gc-community/internal/repositories"
)

type Event interface {
	GetAll() ([]*models.Events, error)
}

type eventUsecase struct {
	er repositories.Event
}

func NewEventUsecase(er repositories.Event) *eventUsecase {
	return &eventUsecase{er}
}

func (eu *eventUsecase) GetAll() ([]*models.Events, error) {
	event, err := eu.er.All()
	if err != nil {
		return event, err
	}

	return event, nil
}