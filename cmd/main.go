package main

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/db"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/handlers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()

	db, err := db.Init(dbUrl)
	if err != nil {
		panic(err)
	}

	handlers.RegisterRoutes(router, db)

	router.Run(":8080")
}
