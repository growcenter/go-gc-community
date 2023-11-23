package repositories

import (
	"fmt"
	"go-gc-community/internal/models"

	"gorm.io/gorm"
)

type Session interface {
	Find(kind string, content interface{}) (*models.Sessions, error)
	All() ([]*models.Sessions, error)
	AllWithFilter(kind string, content interface{}) ([]*models.Sessions, error)
	UpdateFilter(filterColumn string, content interface{}, changesColumn string, changes string) (*gorm.DB)
	Update(session *models.Sessions) (*models.Sessions, error)
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *sessionRepository {
	return &sessionRepository{db}
}

func (sr *sessionRepository) Find(kind string, content interface{}) (*models.Sessions, error) {
	var session *models.Sessions
	column := fmt.Sprintf("%s = ?", kind)
	err := sr.db.Where(column, content).Find(&session).Error
	if err != nil {
		return session, err
	}

	return session, nil
}

func (sr *sessionRepository) All() ([]*models.Sessions, error) {
	var session []*models.Sessions
	err := sr.db.Find(&session).Error
	if err != nil {
		return session, err
	}

	return session, nil
}

func (sr *sessionRepository) AllWithFilter(kind string, content interface{}) ([]*models.Sessions, error) {
	var session []*models.Sessions
	column := fmt.Sprintf("%s = ?", kind)
	err := sr.db.Where(column, content).Find(&session).Error
	if err != nil {
		return session, err
	}

	return session, nil
}

func (sr *sessionRepository) UpdateFilter(filterColumn string, content interface{}, changesColumn string, changes string) (*gorm.DB) {
	session := models.Sessions{}
	/*err := er.db.Model(&event).Where(filterColumn, content).Update(changesColumn, changes).Error
	if err != nil {
		return events, err
	}*/
	
	return sr.db.Model(&session).Where(filterColumn, content).Update(changesColumn, changes)
}

func (sr *sessionRepository) Update(session *models.Sessions) (*models.Sessions, error) {
	err := sr.db.Save(&session).Error
    if err != nil {
        return session, err
    }

    return session, nil
}