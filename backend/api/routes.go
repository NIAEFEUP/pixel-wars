// Package api provides all the router logic needed to NIxel Wars
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"niaefeup/backend-nixel-wars/controller"
)

/*
AddRoutes adds all the API related routes to the engine...
*/
func AddRoutes(engine *gin.Engine) {
	/*
		API related endpoints should be here...
	*/
	apiGroup := engine.Group("/api")
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

	}
}
