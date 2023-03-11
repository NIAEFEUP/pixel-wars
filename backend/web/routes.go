package web

import (
	"github.com/gin-gonic/gin"
	"niaefeup/backend-nixel-wars/controller"
)

func AddRoutes(engine *gin.Engine) {
	engine.GET("/", controller.ShowCanvas)
}