package albums

import (
	"log"
	"strings"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) AddAlbum(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var requestBody struct {
		Album string `json:"album"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
	}

	var desiredAlbumID uint
	h.DB.Raw("SELECT id FROM albums WHERE name = ?", strings.ToLower(requestBody.Album)).Scan(&desiredAlbumID)
	if desiredAlbumID == 0 {
		c.AbortWithStatusJSON(401, gin.H{"error": "Альбом не найден"})
		return
	}

	if result := h.DB.Exec("INSERT INTO user_albums (user_id, album_id) VALUES (?, ?)", claims.ID, desiredAlbumID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось добавить альбом в ваш аккаунт. Скорее всего, он уже добавлен"})
		return
	}

	c.JSON(200, gin.H{"success": "Альбом успешно добавлен в ваш аккаунт"})
}
