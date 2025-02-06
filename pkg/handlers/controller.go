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

	routes := r.Group("/account")
	routes.Use(middlewares.AuthMiddleware())

	routes.POST("/signup", h.Signup)
	routes.POST("/login", h.Login)
}
