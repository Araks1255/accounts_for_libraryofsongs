package handlers

import (
	"log"
	"strings"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) RemoveBand(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var desiredBand struct {
		Band string `json:"band"`
	}

	if err := c.ShouldBindJSON(&desiredBand); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var desiredBandID uint
	h.DB.Raw("SELECT id FROM bands WHERE name = ?", strings.ToLower(desiredBand.Band)).Scan(&desiredBandID)
	if desiredBandID == 0 {
		c.AbortWithStatusJSON(401, gin.H{"error": "Группа не найдена"})
		return
	}

	if result := h.DB.Exec("DELETE FROM user_bands WHERE user_id = ? AND band_id = ?", claims.ID, desiredBandID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось удалить группу. Возможно, она не добавлена в ваш аккаунт"})
		return
	}

	c.JSON(200, gin.H{"success": "Группа успешно удалена из вашего аккаунта"})
}
