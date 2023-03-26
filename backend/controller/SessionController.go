package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SessionEndpoint creates an session if it doesn't already exist to be used on other
// endpoints that might need it.
func SessionEndpoint(ctx *gin.Context) {
	_, err := ctx.Cookie("sessionUUID")
	if err == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{"error": "client already has an session..."})
		return
	}
	ctx.SetSameSite(http.SameSiteStrictMode)
	//the max-age is 400 days max, which is the maximum allowed by chrome.
	ctx.SetCookie("sessionUUID", uuid.NewString(),
		400*24*3600, "/", strings.Split(ctx.Request.Host, ":")[0], false, true)
}
