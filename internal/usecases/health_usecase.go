package usecases

import (
	"go-gc-community/internal/repositories"
)

type Health interface {
	Check() (err error)
}

type healthUsecase struct {
	hr repositories.Health
}

func NewHealthUsecase(hr repositories.Health) *healthUsecase {
	return &healthUsecase{hr}
}

func (hu *healthUsecase) Check() (err error) {
	return hu.hr.Check()
}