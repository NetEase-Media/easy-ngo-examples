package controller

import (
	"net/http"

	"github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/service"
	"github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/vo"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	name := c.Query("name")
	userDto := service.UserInfo(name)
	ret := vo.ConverFromDTO(userDto)
	c.JSON(http.StatusOK, vo.ResponseSuccessWithData(ret))
}
