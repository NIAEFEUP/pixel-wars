package web

import (
	"niaefeup/backend-nixel-wars/controller"

	"github.com/gin-gonic/gin"
)

// AddRoutes adds all routes related to the api backend of this project
func AddRoutes(engine *gin.Engine) {
	go controller.RedisSubscriptionHandler()
	engine.GET("/", controller.ShowCanvas)
	engine.GET("/subscribe", controller.SubscriptionEndpoint)
	engine.GET("/getSession", controller.SessionEndpoint)
}
