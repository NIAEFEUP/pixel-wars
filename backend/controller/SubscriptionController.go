// Package controller handles the logic of the endpoints
package controller

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"niaefeup/backend-nixel-wars/model"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"gopkg.in/mgo.v2/bson"
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

var ctxSubscription = context.Background()
var ctx = context.Background()

var globalConfig = model.Configuration{}

// RedisSubscriptionHandler is a goroutine that handles the subscriptions events of redis
// it handles it's own connection because redis on SUBSCRIBE mode can't do other operations
func RedisSubscriptionHandler() {
	sub := redisclientSubscription.Subscribe(ctxSubscription, "changes")
	defer func() {
		if err := sub.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()
	ch := sub.Channel()
	for msg := range ch {
		parsedMsg := model.PixelColorUpdatePubSubMessage{}
		if err := bson.Unmarshal([]byte(msg.Payload), &parsedMsg); err != nil {
			fmt.Printf("RedisSubscriptionHandler error: %s\n", err.Error())
			continue
		}
		connections.Range(func(key, value any) bool {
			if key != parsedMsg.ClientUUID {
				conn := value.(model.Connection)
				conn.SubscribedChannel <- parsedMsg
			}
			return true
		})
	}
}

// RedisCreateBitFieldIfNotExists creates an appropriate bit by using the initial configuration of the program
func RedisCreateBitFieldIfNotExists(config *model.Configuration) {
	globalConfig = *config
	canvasExists, err := redisclient.Exists(ctx, "canvas").Result()
	if err != nil {
		fmt.Printf("err redis: %v\n", err)
	}
	if canvasExists != 1 {
		fmt.Println("Canvas doens't exist... creating a new one...")
		_, err = redisclient.SetBit(ctx, "canvas", int64(config.CanvasHeight*config.CanvasWidth*4-1), 1).Result()
		if err != nil {
			fmt.Printf("err on setting bit: %v\n", err)
		}
	}
}

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
		msgDecoded, err := model.DecodePixelColorUpdateMessage(msg)
		if err != nil {
			fmt.Printf("err: %v Ignoring packet...\n", err)
			continue
		}

		internalMessage := model.PixelColorUpdatePubSubMessage{
			ClientUUID: sessionUUID,
			Message:    msgDecoded,
		}
		encodedMessage, err := bson.Marshal(internalMessage)
		if err != nil {
			fmt.Printf("err: %v Ignoring packet...\n", err)
			continue
		}
		redisclient.Publish(ctx, "changes", encodedMessage)
		//get offset
		offset := (int(internalMessage.Message.PosX) + globalConfig.CanvasWidth*int(internalMessage.Message.PosY)) * 4
		//set proper bits
		for i := 0; i < 4; i++ {
			bit := 0
			if int(internalMessage.Message.Color&(1<<i)) > 0 {
				bit = 1
			}
			redisclient.SetBit(ctx, "canvas",
				int64(offset)+int64(3-i),
				bit,
			)
		}
	}
	close(conn.SubscribedChannel)
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
	for data := range conn.SubscribedChannel {
		buf := new(bytes.Buffer)
		err := model.EncodePixelColorUpdateMessage(buf, data.Message)
		if err != nil {
			fmt.Printf("err on send goroutine: %v Ignoring packet...\n", err)
			continue
		}
		err3 := conn.WebSockerConn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
		if err3 != nil {
			fmt.Printf("err on send goroutine: %v\n", err)
			break
		}
	}
}

// SubscriptionEndpoint initializes the websockets connection if it doesn't exist
func SubscriptionEndpoint(ctx *gin.Context) {
	clientID := ""
	newHeader := http.Header{}
	cookie, err := ctx.Cookie("sessionUUID")
	if err == nil {
		if _, ok := connections.Load(cookie); ok {
			ctx.AbortWithStatusJSON(400, map[string]any{"error": "client already subscribed..."})
			return
		}
		clientID = cookie
	} else {
		ctx.AbortWithStatusJSON(401, map[string]any{"error": "client doesn't have a session..."})
		return

	}

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, newHeader)
	if err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatusJSON(400, map[string]any{"error": "client not using the upgrade token... possibly not using websockets."})
		return
	}

	connection := model.Connection{
		WebSockerConn:     ws,
		SubscribedChannel: make(chan model.PixelColorUpdatePubSubMessage)}

	connections.Store(clientID, connection)

	go connectionReceiveHandler(clientID)
	go connectionSendHandler(clientID)
}
