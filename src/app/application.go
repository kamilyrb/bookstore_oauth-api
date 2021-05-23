package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kamilyrb/bookstore_oauth-api/src/clients/cassandra"
	"github.com/kamilyrb/bookstore_oauth-api/src/domain/access_token"
	"github.com/kamilyrb/bookstore_oauth-api/src/http"
	"github.com/kamilyrb/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
