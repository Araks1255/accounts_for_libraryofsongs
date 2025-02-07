package handlers

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func (h handler) GetGenres(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}

	var genres []string
	h.DB.Raw("SELECT genres.name FROM genres "+
		"INNER JOIN user_genres ON genres.id = user_genres.genre_id "+
		"INNER JOIN users ON user_genres.user_id = users.id "+
		"WHERE users.id = ?", claims.ID).Scan(&genres)

	response := convertToMap(genres)
	c.JSON(200, response)
}
