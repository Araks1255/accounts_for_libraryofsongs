package handlers

import (
	"log"
	"strings"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func (h handler) AddSong(c *gin.Context) { // Хэндлер добавления песни к себе
	cookie, err := c.Cookie("token") // Получаем данные под заголовком token из cookie
	if err != nil {                  // Обработка ошибок
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}

	claims, err := utils.ParseToken(cookie) // Парсим данные из токена в переменную claims (объект структуры Claims из models)
	if err != nil {                         // Обработываем ошибки
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}

	var desiredSong struct { // Переменная для хранения песни из запроса
		Song string `json:"song"` // Поле с песней
	}

	if err := c.ShouldBindJSON(&desiredSong); err != nil { // Получение нужной песни из хапроса
		log.Println(err) // Обработка ошибок
		c.AbortWithStatusJSON(401, gin.H{"error": "Песня не найдена"})
		return
	}

	var songID uint                                                                                  // Переменная для хранения айди песни, полученной из запроса
	h.DB.Raw("SELECT id FROM songs WHERE name = ?", strings.ToLower(desiredSong.Song)).Scan(&songID) // Ищем в таблице песен айди песни по названию, сканим в переменную
	if songID == 0 {                                                                                 // Если найденный айди равен 0 (песня не существует)
		c.AbortWithStatusJSON(401, gin.H{"error": "Песня не найдена"}) // Уведомляем об этом
		return
	}

	if result := h.DB.Exec("INSERT INTO user_songs (user_id, song_id) VALUES (?, ?)", claims.ID, songID); result.Error != nil { // Вставляем в таблицу отношений айди юзера из токена и песни из переменной
		log.Println(result.Error) // Обработка ошибок
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось добавить песню"})
		return
	}

	c.JSON(200, gin.H{"success": "Песня успешно добавлена"}) // Если нигде не возникло ошибок, то высылаем уведомление об успехе операции
}
