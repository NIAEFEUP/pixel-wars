// main package of NIxel-Wars backend
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"niaefeup/backend-nixel-wars/api"
	"niaefeup/backend-nixel-wars/controller"
	"niaefeup/backend-nixel-wars/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func addReverseProxyPath(path string, r *gin.Engine, destURL *url.URL) {
	r.Any(path, func(ctx *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(destURL)
		proxy.Director = func(r *http.Request) {
			r.Header = ctx.Request.Header
			r.Host = destURL.Host
			r.URL.Scheme = destURL.Scheme
			r.URL.Host = destURL.Host
			r.URL.Path = strings.Split(path, "/*path")[0] + ctx.Param("path")
			r.URL.RawQuery = ctx.Request.URL.RawQuery
		}
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})
}

// AddFrontendRoutes adds the front dynamically depending if it's running on debugmode (forward to vite)
// or on deploy mode
func AddFrontendRoutes(r *gin.Engine, config *model.Configuration) {
	if config.DebugMode {
		remote, err := url.Parse("http://localhost:5173")
		if err != nil {
			panic(err)
		}
		addReverseProxyPath("/", r, remote)
		addReverseProxyPath("/assets/*path", r, remote)
		addReverseProxyPath("/favicons/*path", r, remote)
		addReverseProxyPath("/@vite/*path", r, remote)
		addReverseProxyPath("/src/*path", r, remote)
		addReverseProxyPath("/node_modules/*path", r, remote)

	} else {
		r.Static("/assets", "../frontend/dist/assets")
		r.Static("/favicons", "../frontend/dist/favicons")
		r.StaticFile("/", "../frontend/dist/index.html")
	}

}

func main() {
	config := model.LoadConfigurationFile()
	controller.RedisCreateBitFieldIfNotExists(&config)

	r := gin.Default()
	r.Use(cors.Default())

	api.AddRoutes(r)
	AddFrontendRoutes(r, &config)

	//TODO: serve this as HTTPS
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server...")
		fmt.Println(err.Error())
	}
}
