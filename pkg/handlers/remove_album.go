package handlers

import (
	"log"
	"strings"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) RemoveAlbum(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var requestBody struct {
		Album string `json:"album"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var desiredAlbumID uint
	h.DB.Raw("SELECT id FROM albums WHERE name = ?", strings.ToLower(requestBody.Album)).Scan(&desiredAlbumID)
	if desiredAlbumID == 0 {
		c.AbortWithStatusJSON(401, gin.H{"error": "Альбом не найден"})
		return
	}

	if result := h.DB.Exec("DELETE FROM user_albums WHERE user_id = ? AND album_id = ?", claims.ID, desiredAlbumID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось удалить альбом. Возможно он не добавлен в ваш аккаунт"})
		return
	}

	c.JSON(200, gin.H{"succes": "Альбом успешно удалён из вашего аккаунта"})
}
