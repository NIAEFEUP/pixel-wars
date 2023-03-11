package controller

import (
	"github.com/gin-gonic/gin"

	"niaefeup/backend-nixel-wars/model"
	)

func ShowCanvas(ctx *gin.Context) {
	var canvas = model.Canvas{Valid: true} // change this later
	ctx.JSON(200, gin.H{"colors": canvas.Colors})
}