package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h handler) GetGenres(c *gin.Context) {
	claims, err := ParseClaims(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}

	var genres []string
	h.DB.Raw("SELECT genres.name FROM genres "+
		"INNER JOIN user_genres ON genres.id = user_genres.genre_id "+
		"INNER JOIN users ON user_genres.user_id = users.id "+
		"WHERE users.id = ?", claims.ID).Scan(&genres)

	response := ConvertToMap(genres)
	c.JSON(200, response)
}
