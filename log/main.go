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
	"github.com/NetEase-Media/easy-ngo/application/r/rlog"
	_ "github.com/NetEase-Media/easy-ngo/application/r/rlog/rnlog"
	_ "github.com/NetEase-Media/easy-ngo/application/r/rlog/rzap"
	"github.com/NetEase-Media/easy-ngo/xlog/nlog"
	"github.com/NetEase-Media/easy-ngo/xlog/xfmt"
)

// go run main.go -c ./app.yaml
func main() {
	// Option Nlog
	app := application.Default()
	app.Initialize()
	app.Startup()
	nlogLogger := rlog.GetLogger("nlog2")
	if nlogLogger == nil {
		fmt.Print("failed....")
		return
	}
	nlogLogger.Infof("hello world1")

	// Option Xzap
	xzapLogger := rlog.GetLogger("xzap1")
	if xzapLogger == nil {
		fmt.Print("failed....")
		return
	}
	xzapLogger.Infof("hello world2")

	// Default Xfmt
	xfmt.Infof(xfmt.String())
	xfmt.Infof("hello world11")
	xfmt.Debugf("hello world22")
	xfmt.Infof(xfmt.GetLevel())
	xfmt.SetLevel("debug")
	xfmt.Infof(xfmt.String())
	xfmt.Debugf("hello world33")

	// Default Nlog
	nlog.Infof(nlog.String())
	nlog.Infof("hello world111")
	nlog.Debugf("hello world222")
	nlog.Infof(nlog.GetLevel())
	nlog.SetLevel("debug")
	nlog.Debugf("hello world333")
	nlog.SetFlags("Ldate | Ltime ")
	nlog.Infof(nlog.String())
	nlog.Debugf("hello world444")
}
