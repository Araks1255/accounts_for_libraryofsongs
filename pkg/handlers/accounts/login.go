package accounts

import (
	"log"
	"time"

	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h handler) Login(c *gin.Context) {
	viper.SetConfigFile("./pkg/common/envs/.env") // Настройка вайпера для получен
	viper.ReadInConfig()

	secretKey := []byte(viper.Get("SECRET_KEY").(string))

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	h.DB.Raw("SELECT * FROM users WHERE name = ?", user.Name).Scan(&existingUser)
	if existingUser.ID == 0 {
		log.Println("Пользователь не найден")
		c.AbortWithStatusJSON(401, gin.H{"error": "Пользователь не найден"})
		return
	}

	if ok := utils.CompareHashAndPassword(user.Password, existingUser.Password); !ok {
		log.Println("Неверный пароль")
		c.AbortWithStatusJSON(401, gin.H{"error": "Неверный пароль"})
		return
	}

	expirationTime := time.Now().Add(744 * time.Hour)

	claims := models.Claims{
		ID: existingUser.ID,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Name,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Не удалось сгенерировать токен"})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

	c.JSON(200, gin.H{"success": "Вы успешно вошли в аккаунт"})
}
