// main package of NIxel-Wars backend
package main

import (
	"fmt"
	"niaefeup/backend-nixel-wars/web"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	/*
		Add your groups here...
	*/
	web.AddRoutes(r)
	//TODO: serve this as HTTPS
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server...")
		fmt.Println(err.Error())
	}
}
