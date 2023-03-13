// Package controller handles the logic of the endpoints
package controller

import (
	"fmt"
	"net/http"
	"niaefeup/backend-nixel-wars/model"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var upgrader = websocket.Upgrader{}

// connections is a concurrent map that gets all connections on this server
var connections = sync.Map{}

var redisclient = redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
	DB:   0,
})

var redisclientSubscription = redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
	DB:   0,
})

func connectionReceiveHandler(sessionUUID string) {
	c, ok := connections.Load(sessionUUID)
	if !ok {
		fmt.Printf("Recv goroutine couldn't find session %s...", c)
		return
	}
	conn := c.(model.Connection)
	for {
		_, msg, err := conn.WebSockerConn.ReadMessage()
		if err != nil {
			fmt.Printf("err on recv goroutine: %v\n", err)
			break
		}
		//send logic goes here
		fmt.Printf("msg: %v\n", msg)
	}
	err := conn.WebSockerConn.Close()
	if err != nil {
		fmt.Printf("err on recv goroutine: %v\n", err)
	}
	connections.Delete(sessionUUID)
}

func connectionSendHandler(sessionUUID string) {
	//TODO: do the subscription loader
	c, ok := connections.Load(sessionUUID)
	if !ok {
		fmt.Printf("Recv goroutine couldn't find session %s...", c)
		return
	}
	conn := c.(model.Connection)
	for {
		data := <-conn.SubscribedChannel
		err := conn.WebSockerConn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			fmt.Printf("err on send goroutine: %v\n", err)
			break
		}
		fmt.Printf("msg: %s", data)
	}
	err := conn.WebSockerConn.Close()
	if err != nil {
		fmt.Printf("err on recv goroutine: %v\n", err)
	}
	connections.Delete(sessionUUID)
}

// SubscriptionEndpoint initializes the websockets connection if it doesn't exist
func SubscriptionEndpoint(ctx *gin.Context) {
	clientID := ""
	newHeader := http.Header{}
	cookie, err := ctx.Cookie("sessionUUID")
	if err == nil {
		if _, ok := connections.Load(cookie); ok {
			ctx.AbortWithStatusJSON(400, map[string]any{"error": "client already subscribed"})
			return
		}
		clientID = cookie
	} else {
		newUUID := uuid.NewString()
		newHeader := make(http.Header)
		newHeader.Set("Set-Cookies", fmt.Sprintf("sessionUUID=%s", newUUID))

	}

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, newHeader)
	if err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatusJSON(400, map[string]any{"error": "client not using the upgrade token... possibly not using websockets."})
		return
	}

	connection := model.Connection{
		WebSockerConn:     ws,
		SubscribedChannel: make(chan []uint8)}

	connections.Store(clientID, connection)

	go connectionReceiveHandler(clientID)
	go connectionSendHandler(clientID)
}
