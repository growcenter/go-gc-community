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
	FindMultipleExact(firstParam string, secondParam string, input string) (*models.Registrations, error)
	Update(reg *models.Registrations) (*models.Registrations, error)
	List(page, limit, sort, filter string) ([]*models.Registrations, error)
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

func (rr *registrationRepository) FindMultipleExact(firstParam string, secondParam string, input string) (*models.Registrations, error) {
	var reg *models.Registrations
	column := fmt.Sprintf("%s = ? OR %s = ?", firstParam, secondParam)
	err := rr.db.Where(column, input, input).Find(&reg).Error
	if err != nil {
		return reg, err
	}

	return reg, nil
}

func (rr *registrationRepository) Update(reg *models.Registrations) (*models.Registrations, error) {
	err := rr.db.Save(&reg).Error
    if err != nil {
        return reg, err
    }

    return reg, nil
}

func (rr *registrationRepository) List(page, limit, sort, filter string) ([]*models.Registrations, error) {
	var reg []*models.Registrations
	err := rr.db.Scopes(Paginate(page, limit), Sort(sort), TripleFilter("name LIKE ? OR identifier LIKE ? OR account_number LIKE ?", filter)).Find(&reg).Error
	if err != nil {
		return reg, err
	}

	return reg, nil
}