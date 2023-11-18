package usecases

import (
	"go-gc-community/internal/repositories"

	"go-gc-community/pkg/authorization"
	"go-gc-community/pkg/google"
)

type Dependencies struct {
	Repository		*repositories.Repositories
	Authorization	*authorization.Auth
	Google			*google.Goog
}

type Usecases struct {
	Health	healthUsecase
	User	userUsecase
	Event 	eventUsecase
}

func NewUsecases(dep Dependencies) *Usecases {
	health := NewHealthUsecase(dep.Repository.Health)
	user := NewUserUsecase(dep.Repository.User, dep.Authorization, *dep.Google)
	event := NewEventUsecase(dep.Repository.Event)

	return &Usecases{
		Health: *health,
		User: *user,
		Event: *event,
	}
}