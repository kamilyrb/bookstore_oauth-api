package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kamilyrb/bookstore_oauth-api/src/http"
	"github.com/kamilyrb/bookstore_oauth-api/src/repository/db"
	"github.com/kamilyrb/bookstore_oauth-api/src/repository/rest"
	"github.com/kamilyrb/bookstore_oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {

	atService := access_token.NewService(rest.NewRepository(), db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8001")
}
