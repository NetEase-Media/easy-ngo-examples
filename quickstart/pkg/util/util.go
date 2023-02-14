package util

import (
	"fmt"

	"github.com/NetEase-Media/easy-ngo-examples/quickstart/pkg/constant"
)

func UserRedisKey(name string) string {
	return fmt.Sprintf(constant.UserKeyPrefix, name)
}
