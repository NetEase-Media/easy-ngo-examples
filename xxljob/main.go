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

	"github.com/NetEase-Media/easy-ngo/application"
	_ "github.com/NetEase-Media/easy-ngo/application/r/rconfig"
	"github.com/NetEase-Media/easy-ngo/application/r/rxxljob"
	"github.com/xxl-job/xxl-job-executor-go"
)

func main() {
	app := application.Default()
	app.Initialize()
	app.Startup()

	manager := rxxljob.GetXJobManager("ngo.client.xxxljob")
	if manager == nil {
		fmt.Print("init failed.")
		return
	}

	manager.RegTask("demoJobHandler", Job01)
	// time.Sleep(time.Minute) // 等待调度
	err := manager.Run()
	if err != nil {
		panic("run exception")
	}
	var exit chan int32
	<-exit
}

func Job01(cxt context.Context, param *xxl.RunReq) (msg string) {
	fmt.Println("test one task:" + param.ExecutorHandler + " param:" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	return "test done"
}
