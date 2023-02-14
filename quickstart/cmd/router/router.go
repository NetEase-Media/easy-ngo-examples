package router

import (
	"github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/controller"
	"github.com/NetEase-Media/easy-ngo/application/r/rgin"
)

func Router() error {
	g := rgin.Gin()
	g.POST("/user/auth", controller.UserAuth)
	g.GET("/user/info", controller.UserInfo)
	return nil
}
