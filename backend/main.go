// main package of NIxel-Wars backend
package main

import (
	"fmt"

	"niaefeup/backend-nixel-wars/api"
	"niaefeup/backend-nixel-wars/controller"
	"niaefeup/backend-nixel-wars/model"

	"github.com/gin-gonic/gin"
)

func AddFrontendRoutes(r *gin.Engine) {
	r.Static("/assets", "../frontend/dist/assets")
	r.Static("/favicons", "../frontend/dist/favicons")
	r.StaticFile("/", "../frontend/dist/index.html")
}

func main() {
	r := gin.Default()

	config := model.LoadConfigurationFile()
	controller.RedisCreateBitFieldIfNotExists(&config)

	api.AddRoutes(r)

	AddFrontendRoutes(r)

	//TODO: serve this as HTTPS
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server...")
		fmt.Println(err.Error())
	}
}
