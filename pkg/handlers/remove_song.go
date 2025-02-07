package handlers

import (
	"log"

	//"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func (h handler) RemoveSong(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}
	// Этот кусок надо будет в функцию вынести
	claims, err := utils.ParseToken(cookie)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}

	var desiredSong struct {
		Song string `json:"song"`
	}

	if err := c.ShouldBindJSON(&desiredSong); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var desiredSongID uint

	h.DB.Raw("SELECT id FROM songs WHERE name = ?", desiredSong.Song).Scan(&desiredSongID)
	if desiredSongID == 0 {
		c.AbortWithStatusJSON(401, gin.H{"error": "Песня не найдена"})
		return
	}

	if result := h.DB.Exec("DELETE FROM user_songs WHERE user_id = ? AND song_id = ?", claims.ID, desiredSongID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось удалить песню"})
		return
	}

	c.JSON(200, gin.H{"success": "Песня успешно удалена"})
}
