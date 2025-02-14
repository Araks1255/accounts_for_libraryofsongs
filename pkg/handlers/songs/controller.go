package songs

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/middlewares"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	secretKey := viper.Get("SECRET_KEY").(string)

	h := handler{
		DB: db,
	}

	accounts := r.Group("/account")
	accounts.Use(middlewares.AuthMiddleware(secretKey))

	accounts.POST("/create-song", h.CreateSong)

	songs := accounts.Group("/songs")
	songs.GET("/", h.GetSongs)
	songs.POST("/", h.AddSong)
	songs.DELETE("/", h.RemoveSong)
}
