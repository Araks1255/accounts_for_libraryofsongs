package songs

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := handler{
		DB: db,
	}

	accounts := r.Group("/account")
	accounts.Use(middlewares.AuthMiddleware())

	accounts.POST("/create-song", h.CreateSong)

	songs := accounts.Group("/songs")
	songs.GET("/", h.GetSongs)
	songs.POST("/", h.AddSong)
	songs.DELETE("/", h.RemoveSong)
}
