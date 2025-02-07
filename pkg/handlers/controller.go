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

	// Create song
	accounts.POST("/create-song", h.CreateSong)

	// User songs
	accounts.GET("/songs", h.GetSongs)
	accounts.POST("/songs", h.AddSong)
	accounts.DELETE("/songs", h.RemoveSong)

	// User genres
	accounts.POST("/genres", h.AddGenre)
	accounts.GET("/genres", h.GetGenres)
	accounts.DELETE("/genres", h.RemoveGenre)
}
