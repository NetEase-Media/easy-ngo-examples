// Copyright 2022 NetEase Media Technology（Beijing）Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/NetEase-Media/easy-ngo/application"
	_ "github.com/NetEase-Media/easy-ngo/application/r/rconfig"
	"github.com/NetEase-Media/easy-ngo/application/r/rredis"
	"github.com/NetEase-Media/easy-ngo/clients/xredis"
)

var (
	client xredis.Redis
)

// go run . -c ./app.yaml
func main() {
	app := application.Default()
	app.Initialize()
	app.Startup()

	key := "key"
	value := "value"

	client = rredis.Get("ngo.client.redis.demo01")
	_, err := client.Set(context.Background(), key, value, time.Second*60).Result()
	if err != nil {
		fmt.Print(err)
	}
	res, err := client.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(res)
}
