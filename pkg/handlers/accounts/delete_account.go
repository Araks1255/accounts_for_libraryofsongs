package accounts

import (
	"log"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteAccount(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)

	var userPassword string
	h.DB.Raw("SELECT password FROM users WHERE id = ?", claims.ID).Scan(&userPassword)
	if userPassword == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Пользователь не найден. Попробуйте перезайти в аккаунт"})
		return
	}

	var requestBody struct {
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	if ok := utils.CompareHashAndPassword(requestBody.Password, userPassword); !ok {
		c.AbortWithStatusJSON(401, gin.H{"error": "Неверный пароль"})
		return
	}

	if result := h.DB.Exec("DELETE FROM users WHERE id = ?", claims.ID); result.Error != nil {
		log.Println(result.Error)
		c.AbortWithStatusJSON(401, gin.H{"error": "Не удалось удалить аккаунт"})
		return
	}

	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	c.JSON(200, gin.H{"success": "Аккаунт успешно удалён"})
}
