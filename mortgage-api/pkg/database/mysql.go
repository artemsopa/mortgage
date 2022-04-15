package database

import (
	"fmt"
	"log"

	"github.com/artomsopun/mortgage/mortgage-api/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(user, password, host, port, name string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, name),
	), &gorm.Config{})

	if err != nil {
		log.Panicln(err)
	}

	if err := DB.AutoMigrate(
		&domain.User{}, &domain.Session{},
	); err != nil {
		log.Panicln(err)
	}

	return DB
}

func GetInstanceDB() *gorm.DB {
	return DB
}
