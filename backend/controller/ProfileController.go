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
	profileBytes, err := profileCmd.Bytes()
	if err == nil {
		fmt.Printf("err: %v\n", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	client := model.Client{}
	if err := json.Unmarshal(profileBytes, &client); err != nil {
		fmt.Printf("err: %v\n", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if client.Profile != nil {
		fmt.Printf("User with UUID %s already has an profile", session)
		ctx.AbortWithStatus(http.StatusNotAcceptable)
		return
	}
	if err := ctx.BindJSON(&client.Profile); err != nil {
		fmt.Println("Couldn't bind profile JSON...")
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	clientJSON, err := json.Marshal(client)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	redisclient.Set(ctx, session, clientJSON, 0)
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
	client := model.Client{}
	if err := json.Unmarshal(profileBytes, &client); err != nil {
		fmt.Printf("err: %v\n", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if client.Profile == nil {
		fmt.Println("Client doesn't have a profile...")
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, client.Profile)
}
