// Package api provides all the router logic needed to NIxel Wars
package api

import (
	"github.com/gin-gonic/gin"
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
			ctx.JSON(200, map[string]any{"api": "test!"})
		})
	}
}
