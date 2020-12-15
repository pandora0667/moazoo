package server

import (
	config "../config"
	"github.com/swaggo/gin-swagger/example/basic/docs"
	"log"
)

func Start() {

	log.Println("System Initialize Start")
	restApi()

}

func restApi() {

	// programatically set swagger info
	docs.SwaggerInfo.Title = "wisoft moazoo"
	docs.SwaggerInfo.Description = "collect data api server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "wisoft.io"
	docs.SwaggerInfo.BasePath = "/v2"

	r := setupRouter()
	r.Run(config.HTTP)

}
