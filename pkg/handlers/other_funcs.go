package handlers

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/utils"

	"github.com/gin-gonic/gin"
)

func ParseClaims(c *gin.Context) (claims *models.Claims, err error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return nil, err
	}

	claims, err = utils.ParseToken(cookie)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func ConvertToMap(slice []string) map[int]string { // Функция преобразования слайса в мапу
	result := make(map[int]string)    // Инициализируем мапу
	for i := 0; i < len(slice); i++ { // В цикле
		result[i+1] = slice[i] // Назначем ключу мапы i+1 значение слайса i
	} // И в итоге получается мапа с порядковыми номерами всех значений
	return result // Возвращаем мапу
}
