package main

import (
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/common/db"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/handlers/accounts"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/handlers/albums"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/handlers/bands"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/handlers/genres"
	"github.com/Araks1255/accounts_for_libraryofsongs/pkg/handlers/songs"

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

	accounts.RegisterRoutes(router, db)
	songs.RegisterRoutes(router, db)
	genres.RegisterRoutes(router, db)
	bands.RegisterRoutes(router, db)
	albums.RegisterRoutes(router, db)

	router.Run(":80")
}
