package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"niaefeup/backend-nixel-wars/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SessionEndpoint creates an session if it doesn't already exist to be used on other
// endpoints that might need it.
func SessionEndpoint(c *gin.Context) {
	_, err := c.Cookie("sessionUUID")
	if err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{"error": "client already has an session..."})
		return
	}
	c.SetSameSite(http.SameSiteStrictMode)
	//the max-age is 400 days max, which is the maximum allowed by chrome.
	newUUID := uuid.NewString()
	c.SetCookie("sessionUUID", newUUID,
		400*24*3600, "/", strings.Split(c.Request.Host, ":")[0], false, true)

	client := model.Client{
		Profile:         nil,
		LastTimestamp:   uint64(time.Now().Unix()),
		RemainingPixels: uint64(globalConfig.PixelsPerMinute),
	}
	clientJSON, err := json.Marshal(&client)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	redisclient.Set(ctx, newUUID, clientJSON, 0)
}
