package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBands(c *gin.Context) {
	claims, err := ParseClaims(c)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}

	var bands []string
	h.DB.Raw("SELECT bands.name FROM bands "+
		"INNER JOIN user_bands ON bands.id = user_bands.band_id "+
		"INNER JOIN users ON user_bands.user_id = users.id "+
		"WHERE users.id = ?", claims.ID).Scan(&bands)

	response := ConvertToMap(bands)
	c.JSON(200, response)
}
