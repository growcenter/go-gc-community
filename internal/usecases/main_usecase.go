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
	Salt			[]byte
}

type Usecases struct {
	Health	healthUsecase
	User	userUsecase
	Event 	eventUsecase
	Internal internalUsecase
}

func NewUsecases(dep Dependencies) *Usecases {
	health := NewHealthUsecase(dep.Repository.Health)
	user := NewUserUsecase(dep.Repository.User, dep.Authorization, *dep.Google, dep.Salt)
	event := NewEventUsecase(dep.Repository.User, dep.Repository.Event, dep.Repository.Session, dep.Repository.Registration)
	internal := NewInternalUsecase(dep.Repository.User, dep.Repository.Event, dep.Repository.Session, dep.Repository.Registration)

	return &Usecases{
		Health: *health,
		User: *user,
		Event: *event,
		Internal: *internal,
	}
}