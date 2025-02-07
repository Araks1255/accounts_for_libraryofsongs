package models

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"gorm.io/gorm"
)

type Song struct { // Структура песни (по ней будет создаваться таблица при автомиграции)
	gorm.Model              // Указание того, что это горм модель
	Name       string       `gorm:"unique"` // Название песни, уникальное
	AlbumID    uint         // Поле для хранения айди альбома, к которому относится таблица
	Album      models.Album `gorm:"foreignKey:AlbumID;references:id"` // Указание отношения этой таблицы к таблице albums по внешнему ключу album_id, ссылающемуся на столбец id в таблице albums
	UserID     uint
	User       User `gorm:"foreignKey:UserID;references:id"`
	// Столбец id создаётся автоматически
}
