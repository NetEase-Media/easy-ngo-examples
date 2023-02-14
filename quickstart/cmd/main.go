package main

import (
	"fmt"

	_ "github.com/NetEase-Media/easy-ngo/application/r/rconfig"

	"github.com/NetEase-Media/easy-ngo-examples/quickstart/cmd/router"
	"github.com/NetEase-Media/easy-ngo/application"
)

func main() {
	app := application.Default()
	app.Initialize(router.Router)
	app.Startup()
	fmt.Print("Hello, World.")
}
