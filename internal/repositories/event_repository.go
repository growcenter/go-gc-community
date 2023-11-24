package repositories

import (
	"fmt"
	"go-gc-community/internal/models"

	"gorm.io/gorm"
)

type Event interface {
	Find(kind string, content interface{}) (*models.Events, error)
	All() ([]*models.Events, error)
	UpdateFilter(filterColumn string, content interface{}, changesColumn string, changes string) (*gorm.DB)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *eventRepository {
	return &eventRepository{db}
}

func (er *eventRepository) Find(kind string, content interface{}) (*models.Events, error) {
	var event *models.Events
	column := fmt.Sprintf("%s = ?", kind)
	err := er.db.Where(column, content).Find(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (er *eventRepository) All() ([]*models.Events, error) {
	var event []*models.Events
	err := er.db.Find(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (er *eventRepository) UpdateFilter(filterColumn string, content interface{}, changesColumn string, changes string) (*gorm.DB) {
	events := models.Events{}
	/*err := er.db.Model(&event).Where(filterColumn, content).Update(changesColumn, changes).Error
	if err != nil {
		return events, err
	}*/

	return er.db.Model(&events).Where(filterColumn, content).Update(changesColumn, changes)
}