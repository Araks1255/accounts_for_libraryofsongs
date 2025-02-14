package accounts

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

	r.POST("/signup", h.Signup)
	r.POST("/login", h.Login)
	r.POST("/logout", h.Logout)

	accounts := r.Group("/account")
	accounts.Use(middlewares.AuthMiddleware(secretKey))

	accounts.DELETE("/", h.DeleteAccount)
}
