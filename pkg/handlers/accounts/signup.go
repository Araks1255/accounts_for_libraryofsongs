package accounts

import (
	"log"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func (h handler) Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var existingUserID uint
	h.DB.Raw("SELECT id FROM users WHERE name = ?", user.Name).Scan(&existingUserID)
	if existingUserID != 0 {
		log.Println("Пользователь уже существует")
		c.AbortWithStatusJSON(401, gin.H{"error": "Пользователь уже существует"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)
	if errHash != nil {
		log.Println("Не получилось сгенерировать хэш пароля")
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось сгенерировать хэш пароля"})
		return
	}

	h.DB.Create(&user)

	c.JSON(200, gin.H{"success": "Регистрация прошла успешно"})
}
