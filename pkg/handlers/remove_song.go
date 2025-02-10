package handlers

import (
	"log"
	"strings"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) RemoveSong(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var requestBody struct { // Переменная для хранения песни из запроса
		Song string `json:"song"` // Само поле с песней
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil { // Биндим песню из запроса в переменную
		log.Println(err) // Обрабатываем ошибки
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var desiredSongID uint                                                                                  // Переменная для айди песни, которую надо удалить
	h.DB.Raw("SELECT id FROM songs WHERE name = ?", strings.ToLower(requestBody.Song)).Scan(&desiredSongID) // Ищем айди в таблице песен по имени из запроса, сканим в переменную
	if desiredSongID == 0 {                                                                                 // Если найденный айди равен 0
		c.AbortWithStatusJSON(401, gin.H{"error": "Песня не найдена"}) // Отправляем ошибку
		return                                                         // Ведь песни не существует
	}

	if result := h.DB.Exec("DELETE FROM user_songs WHERE user_id = ? AND song_id = ?", claims.ID, desiredSongID); result.Error != nil { // Удаляем из таблицы отношений ряд, в котором айди польщователя равен тому, что в токене, а айди песни тому, что нашли ранее
		log.Println(result.Error) // Обрабатываем ошибки
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось удалить песню. Возможно, она не добавлена в ваш аккаунт"})
		return
	}

	c.JSON(200, gin.H{"success": "Песня успешно удалена из вашего аккаунта"}) // Успех
}
