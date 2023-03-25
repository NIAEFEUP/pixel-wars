package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"niaefeup/backend-nixel-wars/model"

	"github.com/gin-gonic/gin"
)

// AddProfileEndpoint adds a user profile to a given sessionUUID, if it doesn't exist
func AddProfileEndpoint(ctx *gin.Context) {
	session, err := ctx.Cookie("sessionUUID")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	profileCmd := redisclient.Get(ctx, session)
	if _, err := profileCmd.Bytes(); err == nil {
		fmt.Printf("User with UUID %s already has an profile", session)
		ctx.AbortWithStatus(http.StatusNotAcceptable)
		return
	}
	profile := model.Profile{}
	if err := ctx.BindJSON(&profile); err != nil {
		fmt.Printf("err: %v\n", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	profileJSON, err := json.Marshal(profile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	redisclient.Set(ctx, session, profileJSON, 0)
	ctx.AbortWithStatus(http.StatusOK)
}

// GetProfileEndpoint gets the current profile according to the session cookie.
func GetProfileEndpoint(ctx *gin.Context) {
	session, err := ctx.Cookie("sessionUUID")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	profileCmd := redisclient.Get(ctx, session)
	profileBytes, err := profileCmd.Bytes()
	if err != nil {
		fmt.Printf("Redis get err: %v\n", profileCmd.Err())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	profile := model.Profile{}
	if err := json.Unmarshal(profileBytes, &profile); err != nil {
		fmt.Printf("err: %v\n", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, profile)
}
