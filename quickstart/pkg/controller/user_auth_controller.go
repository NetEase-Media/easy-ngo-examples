package controller

import (
	"net/http"

	"github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/service"
	"github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/vo"
	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	p := vo.AuthParam{}
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(http.StatusUnauthorized, vo.ResponseFailed())
		return
	}
	success := service.Auth(p.Name, p.Passport)
	if success {
		c.JSON(http.StatusOK, vo.ResponseSuccess())
	} else {
		c.JSON(http.StatusUnauthorized, vo.ResponseFailed())
	}
}
