package dao

import (
	"fmt"

	"github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/do"
	"github.com/NetEase-Media/easy-ngo/application/r/rgorm"
)

func GetUser(name string) *do.UserInfo {
	cli := rgorm.GetDBClient("ngo.client.gorm.demo01")
	if cli == nil {
		fmt.Print("failed.")
		return nil
	}
	var user = new(do.UserInfo)
	err := cli.Where("name", name).Find(user).Error
	if err != nil {
		return nil
	}
	return user
}

func UserPassport(name string) string {
	cli := rgorm.GetDBClient("ngo.client.gorm.demo01")
	if cli == nil {
		return ""
	}
	var user = new(do.UserInfo)
	err := cli.Where("name", name).Find(user).Error
	if err != nil {
		return ""
	}
	return user.Passport
}
