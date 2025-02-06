package db

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbUrl string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
