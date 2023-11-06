package repositories

import "gorm.io/gorm"

type Repositories struct {
	Health Health
}

func NewRepositories (db *gorm.DB) *Repositories {
	return &Repositories{
		Health: NewHealthRepository(db),
	}
}