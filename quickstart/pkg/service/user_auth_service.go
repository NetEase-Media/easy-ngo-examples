package service

import "github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/dao"

func Auth(name, passport string) bool {
	pas := dao.UserPassport(name)
	return pas == passport
}
