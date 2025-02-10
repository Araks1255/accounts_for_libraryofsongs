package handlers

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

	r.POST("/signup", h.Signup)
	r.POST("/login", h.Login)
	r.POST("/logout", h.Logout)

	accounts := r.Group("/account")
	accounts.Use(middlewares.AuthMiddleware())

	accounts.DELETE("/", h.DeleteAccount)

	accounts.POST("/create-song", h.CreateSong)

	songs := accounts.Group("/songs")
	songs.GET("/", h.GetSongs)
	songs.POST("/", h.AddSong)
	songs.DELETE("/", h.RemoveSong)

	genres := accounts.Group("/genres")
	genres.POST("/", h.AddGenre)
	genres.GET("/", h.GetGenres)
	genres.DELETE("/", h.RemoveGenre)

	bands := accounts.Group("/bands")
	bands.POST("/", h.AddBand)
	bands.GET("/", h.GetBands)
	bands.DELETE("/", h.RemoveBand)

	albums := accounts.Group("/albums")
	albums.POST("/", h.AddAlbum)
	albums.GET("/", h.GetAlbums)
	albums.DELETE("/", h.RemoveAlbum)
}
