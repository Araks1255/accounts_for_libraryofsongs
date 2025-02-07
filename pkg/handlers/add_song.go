package handlers

import (
	"log"
	"strings"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func (h handler) AddSong(c *gin.Context) { // Хэндлер добавления песни к себе
	cookie, err := c.Cookie("token")
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}
	// Тут процесс получения id пользователя из jwt токена, его поясню, когда в отдельную функцию вынесу
	claims, err := utils.ParseToken(cookie)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}

	var desiredSong struct { // Переменная для хранения песни из запроса
		Song string `json:"song"`
	}

	if err := c.ShouldBindJSON(&desiredSong); err != nil { // Получение нужной песни из хапроса
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var songID uint

	h.DB.Raw("SELECT id FROM songs WHERE name = ?", strings.ToLower(desiredSong.Song)).Scan(&songID)
	if songID == 0 {
		c.AbortWithStatusJSON(401, gin.H{"error": "Песня не найдена"})
		return
	}

	if result := h.DB.Exec("INSERT INTO user_songs (user_id, song_id) VALUES (?, ?)", claims.ID, songID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось добавить песню"})
		return
	}

	c.JSON(200, gin.H{"success": "Песня успешно добавлена"})
}
