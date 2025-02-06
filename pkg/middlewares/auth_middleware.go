package middlewares

import (
	"log"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		claims, err := utils.ParseToken(cookie)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(401, gin.H{"error": "unathorized"})
			return
		}

		c.Set("id", claims.ID)
		c.Next()
	}
}
