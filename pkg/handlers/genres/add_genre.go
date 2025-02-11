package genres

import (
	"log"
	"strings"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) AddGenre(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var requestBody struct {
		Genre string `json:"genre"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var genreID uint
	h.DB.Raw("SELECT id FROM genres WHERE name = ?", strings.ToLower(requestBody.Genre)).Scan(&genreID)
	if genreID == 0 {
		c.AbortWithStatusJSON(401, gin.H{"error": "Жанр не найден"})
		return
	}

	if result := h.DB.Exec("INSERT INTO user_genres (user_id, genre_id) VALUES (?, ?)", claims.ID, genreID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось добавить жанр. Скорее всего, он уже добавлен в ваш аккаунт"})
		return
	}

	c.JSON(200, gin.H{"access": "Жанр успешно добавлен"})
}
