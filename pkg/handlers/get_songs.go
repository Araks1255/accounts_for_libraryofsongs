package handlers

import (
	"log"

	//"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func (h handler) GetSongs(c *gin.Context) {
	cookie, err := c.Cookie("token")
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

	var songs []string

	h.DB.Raw("SELECT songs.name FROM songs "+
	"INNER JOIN user_songs ON songs.id = user_songs.song_id "+
	"INNER JOIN users ON user_songs.user_id = users.id "+
	"WHERE users.id = ?", claims.ID).Scan(&songs)

	response := convertToMap(songs)
	c.JSON(200, response)
}

func convertToMap(slice []string) map[int]string {
	result := make(map[int]string)
	for i := 0; i < len(slice); i++ {
		result[i+1] = slice[i]
	}
	return result
}
