package middlewares

import (
	"log"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
			return
		}

		claims, err := utils.ParseToken(cookie, secretKey)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
			return
		}

		c.Set("claims", claims)
	}
}
