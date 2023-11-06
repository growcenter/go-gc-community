package msql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(user, password, host, name string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Database %s is connected.", name)

	return db, nil
}