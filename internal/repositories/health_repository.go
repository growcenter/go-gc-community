package repositories

import "gorm.io/gorm"

type Health interface {
	Check() (err error)
}

type healthRepository struct {
	db *gorm.DB
}

func NewHealthRepository(db *gorm.DB) *healthRepository {
	return &healthRepository{db}
}

func (hr *healthRepository) Check() (err error) {
	//sqlDb, err := hr.db.DB()
	msql, err := hr.db.DB()
	if err != nil {
		return err
	}
	err = msql.Ping()
	if err != nil {
		return err
	}
	
	return nil
}