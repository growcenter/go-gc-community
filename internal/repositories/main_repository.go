package repositories

import "gorm.io/gorm"

type Repositories struct {
	Health Health
	User User
}

func NewRepositories (db *gorm.DB) *Repositories {
	return &Repositories{
		Health: NewHealthRepository(db),
		User: NewUserRepository(db),
	}
}