package vo

import "github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/dto"

type UserInfoVO struct {
	Id      int32  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Gender  int32  `json:"gender"`
}

func ConverFromDTO(d *dto.UserInfoDTO) *UserInfoVO {
	if d == nil {
		return nil
	}
	return &UserInfoVO{
		Id:      d.Id,
		Name:    d.Name,
		Address: d.Address,
		Gender:  d.Gender,
	}
}
