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
	conf "github.com/NetEase-Media/easy-ngo/config"
)

// example01
// go run . -c ./app.toml
// example02
// go run .
// example01 & example02 have the same result.
func main() {
	app := application.Default()
	app.Initialize()
	app.Startup()
	ip := conf.GetString("alpha.ip")
	fmt.Print(ip)
}
