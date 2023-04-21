package config

import (
	"exchange-api/pkg/external/model"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {

	return gorm.Open(sqlite.Open("exchange.db"))
}

func Migrate() {

	db, err := GetConnection()

	if err != nil {
		log.Fatalf("Could not establish database connection, error: %v", err)
		panic(err)
	}

	db.AutoMigrate(&model.Quotation{})
}
