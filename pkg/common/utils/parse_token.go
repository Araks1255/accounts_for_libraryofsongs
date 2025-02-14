package utils

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func ParseToken(tokenString string) (claims *models.Claims, err error) {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	secretKey := viper.Get("SECRET_KEY").(string)

	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), err
	})

	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
