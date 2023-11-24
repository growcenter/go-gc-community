package repositories

import (
	"fmt"
	"go-gc-community/internal/models"

	"gorm.io/gorm"
)

type Registration interface {
	Create(reg *models.Registrations) (*models.Registrations, error)
	BatchCreate(reg []*models.Registrations) ([]*models.Registrations, error)
	Find(kind string, content interface{}) (*models.Registrations, error)
	FindBatchExclude(kind string, content interface{}, kinds string, contents interface{}) ([]*models.Registrations, error)
}

type registrationRepository struct {
	db *gorm.DB
}

func NewRegistrationRepository(db *gorm.DB) *registrationRepository {
	return &registrationRepository{db}
}

func (rr *registrationRepository) Create(reg *models.Registrations) (*models.Registrations, error){
	err := rr.db.Create(&reg).Error
	if err != nil {
		return reg, err
	}

	return reg, nil
}

func (rr *registrationRepository) BatchCreate(reg []*models.Registrations) ([]*models.Registrations, error) {
	err := rr.db.Create(&reg).Error
	if err != nil {
		return reg, err
	}

	return reg, nil
}

func (rr *registrationRepository) Find(kind string, content interface{}) (*models.Registrations, error) {
	var reg *models.Registrations
	column := fmt.Sprintf("%s = ?", kind)
	err := rr.db.Where(column, content).Find(&reg).Error
	if err != nil {
		return reg, err
	}

	return reg, nil
}

func (rr *registrationRepository) FindBatchExclude(kind string, content interface{}, kinds string, contents interface{}) ([]*models.Registrations, error) {
	var reg []*models.Registrations
	column := fmt.Sprintf("%s = ?", kind)
	columns := fmt.Sprintf("%s = ?", kinds)
	err := rr.db.Where(column, content).Not(columns, contents).Find(&reg).Error
	if err != nil {
		return reg, err
	}

	return reg, nil
}