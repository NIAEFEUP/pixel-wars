package web

import (
	"niaefeup/backend-nixel-wars/controller"

	"github.com/gin-gonic/gin"
)

// AddRoutes adds all routes related to the api backend of this project
func AddRoutes(engine *gin.Engine) {
	engine.GET("/", controller.ShowCanvas)
}
