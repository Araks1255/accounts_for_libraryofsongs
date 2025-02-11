package songs

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func (h handler) GetSongs(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var songs []string                        // Переменная для хранения названий найденных песен
	h.DB.Raw("SELECT songs.name FROM songs "+ // Сырой запрос для поиска всех песен юзера
		"INNER JOIN user_songs ON songs.id = user_songs.song_id "+ // Через таблицу отношений
		"INNER JOIN users ON user_songs.user_id = users.id "+
		"WHERE users.id = ?", claims.ID).Scan(&songs) // Сканируем найденные значения в songs

	response := utils.ConvertToMap(songs) // Делаем из массива мапу
	c.JSON(200, response)           // Отправляем её как JSON
}
