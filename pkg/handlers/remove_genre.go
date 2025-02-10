package handlers

import (
	"log"
	"strings"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) RemoveGenre(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var requestBody struct {
		Genre string `json:"genre"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var desiredGenreID uint
	h.DB.Raw("SELECT id FROM genres WHERE name = ?", strings.ToLower(requestBody.Genre)).Scan(&desiredGenreID)
	if desiredGenreID == 0 {
		c.AbortWithStatusJSON(401, gin.H{"error": "Жанр не найден"})
		return
	}

	if result := h.DB.Exec("DELETE FROM user_genres WHERE user_id = ? AND genre_id = ?", claims.ID, desiredGenreID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось удалить жанр. Возможно, он не добавлен в ваш аккаунт"})
		return
	}

	c.JSON(200, gin.H{"success": "Жанр был успешно удалён из вашего аккаунта"})
}
