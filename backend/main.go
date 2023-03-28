// main package of NIxel-Wars backend
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
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
		addReverseProxyPath("/pixelwars/", r, remote)
		addReverseProxyPath("/pixelwars/assets/*path", r, remote)
		addReverseProxyPath("/pixelwars/favicons/*path", r, remote)
		addReverseProxyPath("/pixelwars/@vite/*path", r, remote)
		addReverseProxyPath("/pixelwars/src/*path", r, remote)
		addReverseProxyPath("/pixelwars/node_modules/*path", r, remote)

	} else {
		r.Static("/assets", "../frontend/dist/assets")
		r.Static("/favicons", "../frontend/dist/favicons")
		r.StaticFile("/", "../frontend/dist/index.html")
	}

}

func main() {
	config := model.Configuration{}
	cmdArguments := os.Args[1:]
	if len(cmdArguments) == 1 && cmdArguments[0] == "--prod" {
		config = model.LoadConfigurationFile("./config_prod.json")

	} else if len(cmdArguments) == 0 {
		config = model.LoadConfigurationFile("./config.json")

	} else {
		fmt.Println("Command line argument not recognized...")
		os.Exit(1)
	}
	controller.RedisCreateBitFieldIfNotExists(&config)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowWebSockets: true,
		AllowMethods:    []string{"*"},
		AllowOrigins:    []string{"*"},
	}))

	api.AddRoutes(r, &config)
	AddFrontendRoutes(r, &config)

	if !config.DebugMode {
		if err := r.Run(":80"); err != nil {
			fmt.Println("Failed to start server...")
			fmt.Println(err.Error())
			panic("siuuuu")
		}
	} else {
		if err := r.RunTLS(":8080", "./certs/"+config.Host+".pem", "./certs/"+config.Host+"-key.pem"); err != nil {
			fmt.Println("Failed to start server...")
			fmt.Println(err.Error())
			panic("siuuuu")

		}
	}

}
