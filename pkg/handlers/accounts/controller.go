package accounts

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
}
