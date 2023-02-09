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
	"fmt"

	"github.com/NetEase-Media/easy-ngo/application"
	_ "github.com/NetEase-Media/easy-ngo/application/r/rconfig"

	// _ "github.com/NetEase-Media/easy-ngo/application/r/rgorm"

	"github.com/NetEase-Media/easy-ngo/application/r/rgorm"
)

func main() {
	app := application.Default()
	app.Initialize()
	app.Startup()

	cli := rgorm.GetDBClient("ngo.client.gorm.demo01")
	if cli == nil {
		fmt.Print("failed.")
		return
	}
	err := cli.Exec("update blacklist set blackword = ? where id = ?", "mmm", 10).Error
	if err != nil {
		fmt.Print("err:", err)
		return
	}
	fmt.Print("success")
}
