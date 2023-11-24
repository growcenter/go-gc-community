package repositories

import "gorm.io/gorm"

type Repositories struct {
	Health Health
	User User
	Event Event
	Session Session
	Registration Registration
}

func NewRepositories (db *gorm.DB) *Repositories {
	return &Repositories{
		Health: NewHealthRepository(db),
		User: NewUserRepository(db),
		Event: NewEventRepository(db),
		Session: NewSessionRepository(db),
		Registration: NewRegistrationRepository(db),
	}
}