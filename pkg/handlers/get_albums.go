package handlers

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAlbums(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var albums []string
	h.DB.Raw("SELECT albums.name FROM albums "+
		"INNER JOIN user_albums ON albums.id = user_albums.album_id "+
		"INNER JOIN users ON user_albums.user_id = users.id "+
		"WHERE users.id = ?", claims.ID).Scan(&albums)

	response := ConvertToMap(albums)
	c.JSON(200, response)
}
