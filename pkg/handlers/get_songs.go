package handlers

import (
	"log"

	//"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func (h handler) GetSongs(c *gin.Context) {
	cookie, err := c.Cookie("token") // Пояснение в add_song.go
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Вы не авторизованы"})
		return
	}

	var songs []string                        // Переменная для хранения названий найденных песен
	h.DB.Raw("SELECT songs.name FROM songs "+ // Сырой запрос для поиска всех песен юзера
		"INNER JOIN user_songs ON songs.id = user_songs.song_id "+ // Через таблицу отношений
		"INNER JOIN users ON user_songs.user_id = users.id "+
		"WHERE users.id = ?", claims.ID).Scan(&songs) // Сканируем найденные значения в songs

	response := convertToMap(songs) // Делаем из массива мапу
	c.JSON(200, response)           // Отправляем её как JSON
}

func convertToMap(slice []string) map[int]string { // Функция преобразования слайса в мапу
	result := make(map[int]string)    // Инициализируем мапу
	for i := 0; i < len(slice); i++ { // В цикле
		result[i+1] = slice[i] // Назначем ключу мапы i+1 значение слайса i
	} // И в итоге получается мапа с порядковыми номерами всех значений
	return result // Возвращаем мапу
}
