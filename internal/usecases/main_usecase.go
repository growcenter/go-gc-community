package usecases

import (
	"go-gc-community/internal/repositories"

	"go-gc-community/pkg/google"
	"go-gc-community/pkg/token"
)

type Dependencies struct {
	Repository	*repositories.Repositories
	Token		*token.Auth
	Google		*google.Goog
}

type Usecases struct {
	Health	healthUsecase
	User	userUsecase
}

func NewUsecases(dep Dependencies) *Usecases {
	health := NewHealthUsecase(dep.Repository.Health)
	user := NewUserUsecase(dep.Repository.User, dep.Token, *dep.Google)

	return &Usecases{
		Health: *health,
		User: *user,
	}
}