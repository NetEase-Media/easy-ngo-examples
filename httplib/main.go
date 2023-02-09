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
	"github.com/NetEase-Media/easy-ngo/application/r/rhttplib"
)

// go run . -c ./app.yaml
func main() {

	app := application.Default()
	app.Initialize()
	app.Startup()

	cli := rhttplib.GetHttpClient("ngo.client.http.demo01")
	if cli == nil {
		panic("init failed.")
	}

	code, err := cli.Get("https://www.163.com").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if code != 200 {
		panic("failed.")
	}
	fmt.Print("success")
}
