package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/dao"
	"github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/dto"
	"github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/util"
	"github.com/NetEase-Media/easy-ngo/application/r/rredis"
)

func UserInfo(name string) (ret *dto.UserInfoDTO) {
	redisCli := rredis.Get("ngo.client.redis.demo01")
	key := util.UserRedisKey(name)
	if redisCli != nil {
		bo, err := redisCli.Get(context.Background(), key).Result()
		if err == nil || len(bo) > 0 {
			ret = &dto.UserInfoDTO{}
			err := json.Unmarshal([]byte(bo), ret)
			if err != nil {
				ret = nil
			} else {
				return
			}
		}
	}
	if ret == nil {
		d := dao.GetUser(name)
		if d == nil {
			return nil
		}
		ret = dto.ConvertFromDO(d)
		if ret == nil {
			return
		}
		bo, err := json.Marshal(ret)
		if err != nil {
			return
		}
		redisCli.Set(context.Background(), key, bo, time.Minute)
	}
	return ret
}
