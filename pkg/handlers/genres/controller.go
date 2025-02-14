package genres

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

	genres := accounts.Group("/genres")
	genres.POST("/", h.AddGenre)
	genres.GET("/", h.GetGenres)
	genres.DELETE("/", h.RemoveGenre)
}