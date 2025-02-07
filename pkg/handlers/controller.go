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

	accounts.POST("/song", h.CreateSong)
	accounts.GET("/songs", h.GetSongs)
	accounts.POST("/songs", h.AddSong)
	accounts.DELETE("/songs", h.RemoveSong)

}
