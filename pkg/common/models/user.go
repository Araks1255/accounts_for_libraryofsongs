package models

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string         `gorm:"unique" json:"name"`
	Password string         `json:"password"`
	Songs    []Song         `gorm:"many2many:user_songs;"`
	Albums   []models.Album `gorm:"many2many:user_albums;"`
	Bands    []models.Band  `gorm:"many2many:user_bands;"`
	Genres   []models.Genre `gorm:"many2many:user_genres;"`
}
