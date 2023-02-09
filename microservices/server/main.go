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

	_ "github.com/NetEase-Media/easy-ngo/application/r/rconfig"

	"github.com/NetEase-Media/easy-ngo/application"
	"github.com/NetEase-Media/easy-ngo/application/r/rms/grpc/server"
	"github.com/NetEase-Media/easy-ngo/microservices/testdata"
)

func main() {
	app := application.Default()
	app.Initialize()
	srv := server.Get("server1")
	testdata.RegisterGreeterServer(srv, &GreeterServerImpl{})
	app.Startup()
}

type GreeterServerImpl struct {
	testdata.UnimplementedGreeterServer
}

func (g GreeterServerImpl) SayHello(ctx context.Context, request *testdata.HelloRequest) (*testdata.HelloReply, error) {
	fmt.Println("Hello " + request.Name)
	return &testdata.HelloReply{
		Message: fmt.Sprintf("hello %s", request.Name),
	}, nil
}
