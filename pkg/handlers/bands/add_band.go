package bands

import (
	"log"
	"strings"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) AddBand(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var requestBody struct {
		Band string `json:"band"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var desiredBandID uint
	h.DB.Raw("SELECT id FROM bands WHERE name = ?", strings.ToLower(requestBody.Band)).Scan(&desiredBandID)
	if desiredBandID == 0 {
		c.AbortWithStatusJSON(401, gin.H{"error": "Группа не найдена"})
		return
	}

	if result := h.DB.Exec("INSERT INTO user_bands (user_id, band_id) VALUES (?, ?)", claims.ID, desiredBandID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось добавить группу. Возможно, она уже добавлена в ваш аккаунт"})
		return
	}

	c.JSON(200, gin.H{"success": "Группа успешно добавлена в ваш аккаунт"})
}
