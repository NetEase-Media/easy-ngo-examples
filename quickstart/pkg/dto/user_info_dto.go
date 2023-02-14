package dto

import "github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/do"

type UserInfoDTO struct {
	Id      int32  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Gender  int32  `json:"gender"`
}

func ConvertFromDO(d *do.UserInfo) *UserInfoDTO {
	if d == nil {
		return nil
	}
	return &UserInfoDTO{
		Id:      d.Id,
		Name:    d.Name,
		Address: d.Address,
		Gender:  d.Gender,
	}
}
