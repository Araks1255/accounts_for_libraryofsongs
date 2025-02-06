package handlers

import (
	"log"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	song "github.com/Araks1255/libraryofsongs/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) AddSong(c *gin.Context) {
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

	var desiredSong struct {
		Song string `json:"song"`
	}

	if err := c.ShouldBindJSON(&desiredSong); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var song song.Song

	h.DB.Raw("SELECT * FROM songs WHERE name = ?", desiredSong.Song).Scan(&song)
	if song.ID == 0 {
		c.AbortWithStatusJSON(401, gin.H{"error": "Песня не найдена"})
		return
	}

	var user models.User

	if result := h.DB.Preload("songs").First(&user, claims.ID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(401, gin.H{"error": result.Error.Error()})
		return
	}

	user.Songs = append(user.Songs, song)

	h.DB.Save(&user)

	c.JSON(200, gin.H{"success": "Песня успешно добавлена"})
}
