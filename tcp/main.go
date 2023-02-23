package main

import (
	"context"
	"fmt"
	"net"

	_ "github.com/NetEase-Media/easy-ngo/application/r/rconfig"
	"github.com/NetEase-Media/easy-ngo/application/r/rtcp"

	"github.com/NetEase-Media/easy-ngo/application"
)

func main() {
	app := application.Default()
	app.Initialize()
	server := rtcp.TcpServer()
	server.RegisterHandler(Handler)
	app.Startup()
}

func Handler(conn net.Conn, ctx context.Context) {
	var buf []byte = make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Print(string(buf))
	}
}
