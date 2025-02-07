package handlers

import (
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

	r.POST("/signup", h.Signup)
	r.POST("/login", h.Login)
	r.POST("/logout", h.Logout)

	accounts := r.Group("/account")

	accounts.POST("/create-song", h.CreateSong)

	songs := accounts.Group("/songs")
	songs.GET("/", h.GetSongs)
	songs.POST("/", h.AddSong)
	songs.DELETE("/", h.RemoveSong)

	genres := accounts.Group("/genres")
	genres.POST("/", h.AddGenre)
	genres.GET("/", h.GetGenres)
	genres.DELETE("/", h.RemoveGenre)
}
