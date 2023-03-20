package controller

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"

	"niaefeup/backend-nixel-wars/model"
)

func GetCanvas(ctx *gin.Context) {
	/*
	check if current canvas is valid
		retrieve canvas
	get bitmap from redis
	cache it
	set as valid
	retrieve canvas
	*/
	var canvas model.Canvas = model.Canvas{Valid: true} // change this later
	ctx.JSON(http.StatusOK, gin.H{"colors": canvas.Colors}) // change this later
}

func UpdateCanvas(ctx *gin.Context) {
	i, err := strconv.ParseUint(ctx.Param("offset"), 10, 64)
	i1, err1 := strconv.ParseUint(ctx.Param("color"), 10, 64)
	// Check for errors
	if (err != nil || err1 != nil){
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameters"}) 
	}
	var offset uint32 = uint32(i)
	var color uint8 = uint8(i1)
	ctx.JSON(http.StatusOK, gin.H{"offset": offset, "color": color})
}