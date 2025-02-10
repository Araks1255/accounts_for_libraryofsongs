package handlers

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBands(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var bands []string
	h.DB.Raw("SELECT bands.name FROM bands "+
		"INNER JOIN user_bands ON bands.id = user_bands.band_id "+
		"INNER JOIN users ON user_bands.user_id = users.id "+
		"WHERE users.id = ?", claims.ID).Scan(&bands)

	response := ConvertToMap(bands)
	c.JSON(200, response)
}
