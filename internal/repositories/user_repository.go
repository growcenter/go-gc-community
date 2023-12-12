package repositories

import (
	"fmt"
	"go-gc-community/internal/models"
	"strings"

	"gorm.io/gorm"
)

type User interface {
	Create(user *models.User) (*models.User, error)
	Find(kind string, content string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	First(kind string, content string) (*models.User, error)
	FindMultipleExact(firstParam string, secondParam string, input string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) Create(user *models.User) (*models.User, error){
	err := ur.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepository) Find(kind string, content string) (*models.User, error) {
	var user *models.User
	column := fmt.Sprintf("%s = ?", kind)
	err := ur.db.Where(column, content).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepository) Update(user *models.User) (*models.User, error){
	err := ur.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepository) First(kind string, content string) (*models.User, error) {
	var user *models.User
	column := fmt.Sprintf("%s = ?", kind)
	
	if kind == "email" {
		lowerCase := strings.ToLower(content)
		content = lowerCase
	}
	
	err := ur.db.First(&user, column, content).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepository) FindMultipleExact(firstParam string, secondParam string, input string) (*models.User, error) {
	var user *models.User
	column := fmt.Sprintf("%s = ? OR %s = ?", firstParam, secondParam)
	err := ur.db.Where(column, input, input).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}