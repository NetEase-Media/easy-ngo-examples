package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NetEase-Media/easy-ngo/application"
	_ "github.com/NetEase-Media/easy-ngo/application/r/rconfig"
	"github.com/NetEase-Media/easy-ngo/application/r/rgin"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func WebSocket(c *gin.Context) {
	if websocket.IsWebSocketUpgrade(c.Request) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, c.Writer.Header())
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		ws.WriteMessage(websocket.TextMessage, []byte("Hello Client!"))
		ws.SetCloseHandler(func(code int, text string) error {
			fmt.Println(code, text)
			return nil
		})
		// defer ws.Close()
		go func() {
			for {
				mt, message, err := ws.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					break
				}
				log.Printf("recv: %s", message)
				bytes := []byte("Hello Client!")
				message = append(bytes, message...)
				err = ws.WriteMessage(mt, message)
				if err != nil {
					log.Println("write:", err)
					break
				}
			}
		}()
	}
}
func xgin() error {
	g := rgin.Gin()
	g.GET("/hello", WebSocket)
	return nil
}
func main() {
	app := application.Default()
	app.Initialize(xgin)
	app.Startup()
}
