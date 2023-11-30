package repositories

import (
	"strconv"

	"gorm.io/gorm"
)

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

func Paginate(page, limit string) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        pageNumber, _ := strconv.Atoi(page)
        limitNumber, _ := strconv.Atoi(limit)
        offset := (pageNumber - 1) * limitNumber
        return db.Offset(offset).Limit(limitNumber)
    }
}

func Sort(sort string) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Order(sort)
    }
}

func TripleFilter(key string, filter string) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where(key, "%"+filter+"%", "%"+filter+"%", "%"+filter+"%")
    }
}