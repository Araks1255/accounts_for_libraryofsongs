package genres

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func (h handler) GetGenres(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var genres []string
	h.DB.Raw("SELECT genres.name FROM genres "+
		"INNER JOIN user_genres ON genres.id = user_genres.genre_id "+
		"INNER JOIN users ON user_genres.user_id = users.id "+
		"WHERE users.id = ?", claims.ID).Scan(&genres)

	response := utils.ConvertToMap(genres)
	c.JSON(200, response)
}
