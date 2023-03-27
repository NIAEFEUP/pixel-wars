package controller

import (
	"encoding/base64"
	"net/http"
	"strconv"

	"niaefeup/backend-nixel-wars/model"

	"github.com/gin-gonic/gin"
)

// Canvas is the middle representation of canvas from redis, and that can and probably be will altered through subscription controllers
var Canvas model.Canvas = model.Canvas{Valid: false}

func GetCanvas(ctx *gin.Context) {
	if !Canvas.Valid {
		colors, err := redisclient.Get(ctx, "canvas").Bytes()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving canvas from redis", "details": err.Error()})
		}

		Canvas.Colors = colors
		Canvas.Valid = true
	}

	ctx.JSON(http.StatusOK, gin.H{"canvas": base64.StdEncoding.EncodeToString(Canvas.Colors), "size": gin.H{"height": globalConfig.CanvasHeight, "width": globalConfig.CanvasWidth}})
}

func UpdateCanvas(ctx *gin.Context) {
	i, err := strconv.ParseUint(ctx.Param("offset"), 10, 64)
	i1, err1 := strconv.ParseUint(ctx.Param("color"), 10, 64)
	// Check for errors
	if err != nil || err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameters"})
	}

	Canvas.Valid = false
	var offset uint32 = uint32(i)
	var color uint8 = uint8(i1)
	ctx.JSON(http.StatusOK, gin.H{"offset": offset, "color": color})
}
