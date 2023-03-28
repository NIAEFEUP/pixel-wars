// Package api provides all the router logic needed to NIxel Wars
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"niaefeup/backend-nixel-wars/controller"
	"niaefeup/backend-nixel-wars/model"
)

/*
AddRoutes adds all the API related routes to the engine...
*/
func AddRoutes(engine *gin.Engine, config *model.Configuration) {
	/*
		API related endpoints should be here...
	*/
	groupRoute := ""
	if config.DebugMode {
		groupRoute = "/pixelwars/api"
	} else {
		groupRoute = "/api"

	}

	apiGroup := engine.Group(groupRoute)
	{
		apiGroup.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, map[string]any{"api": "test!"})
		})

		go controller.RedisSubscriptionHandler()
		apiGroup.GET("/subscribe", controller.SubscriptionEndpoint)
		apiGroup.GET("/getSession", controller.SessionEndpoint)

		apiGroup.GET("/canvas", controller.GetCanvas)

		apiGroup.PUT("/canvas/:offset/:color", controller.UpdateCanvas)

		apiGroup.POST("/profiles/new", controller.AddProfileEndpoint)
		apiGroup.GET("/profiles/get", controller.GetProfileEndpoint)
		apiGroup.GET("/client/details", controller.GetClientTimeoutEndpoint)

	}
}
