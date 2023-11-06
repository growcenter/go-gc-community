package usecases

import "go-gc-community/internal/repositories"

type Dependencies struct {
	Repository	*repositories.Repositories	
}

type Usecases struct {
	Health	healthUsecase
}

func NewUsecases(dep Dependencies) *Usecases {
	health := NewHealthUsecase(dep.Repository.Health)

	return &Usecases{
		Health: *health,
	}
}