package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"

	ws "github.com/riete/go-websocket"
)

func main() {
	c, err := ws.NewClient(nil, "ws://localhost:8080/echo", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.SetPingHandler(func(s string) error {
		log.Println("recv ping from server: " + s)
		return nil
	})
	c.SendHeartbeat(context.Background(), time.Second, 3, []byte("hello I'm client"), func(err error) { fmt.Println(1, err) })
	go func() {
		for {
			time.Sleep(time.Second)
			log.Println(c.WriteMessage([]byte("aaa")))
		}
	}()
	c.ReadMessage()
	c.WriteClose(websocket.CloseInternalServerErr, "bbb")
	c.Close()
}
