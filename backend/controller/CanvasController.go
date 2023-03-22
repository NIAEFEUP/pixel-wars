package controller

import (
	"net/http"
	"strconv"

	"niaefeup/backend-nixel-wars/model"

	"github.com/gin-gonic/gin"
)

var canvas model.Canvas = model.Canvas{Valid: false}

func GetCanvas(ctx *gin.Context) {
	if !canvas.Valid {
		colors, err := redisclient.Get(ctx, "canvas").Result()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving canvas from redis", "details": err.Error()})
		}

		canvas.Colors = colors
		canvas.Valid = true
	}

	ctx.JSON(http.StatusOK, canvas.Colors)
}

func UpdateCanvas(ctx *gin.Context) {
	i, err := strconv.ParseUint(ctx.Param("offset"), 10, 64)
	i1, err1 := strconv.ParseUint(ctx.Param("color"), 10, 64)
	// Check for errors
	if err != nil || err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameters"})
	}

	canvas.Valid = false
	var offset uint32 = uint32(i)
	var color uint8 = uint8(i1)
	ctx.JSON(http.StatusOK, gin.H{"offset": offset, "color": color})
}
