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
	"net/http"
	_ "net/http/pprof"
	"runtime"

	_ "github.com/NetEase-Media/easy-ngo/application/r/rconfig"

	_ "github.com/NetEase-Media/easy-ngo/application/r/rotel"

	"github.com/NetEase-Media/easy-ngo/application"

	"github.com/NetEase-Media/easy-ngo/application/r/rgin"
	"github.com/NetEase-Media/easy-ngo/application/r/rgorm"
	"github.com/NetEase-Media/easy-ngo/application/r/rhttplib"
	"github.com/NetEase-Media/easy-ngo/application/r/rkafka"
	"github.com/NetEase-Media/easy-ngo/application/r/rms/grpc/client"
	"github.com/NetEase-Media/easy-ngo/application/r/rms/grpc/server"
	"github.com/NetEase-Media/easy-ngo/application/r/rredis"
	"github.com/NetEase-Media/easy-ngo/microservices/testdata"

	"github.com/gin-gonic/gin"
)

// go run . -c ./app.yaml
func main() {
	//启动pprof
	go func() {
		pprofServer := &http.Server{Addr: "0.0.0.0:19933"}
		pprofServer.ListenAndServe()
	}()

	app := &application.Application{}
	// app初始化完成回调
	app.Initialize(func() error {

		conn := client.Get("rpc-client-1")
		greeterClient := testdata.NewGreeterClient(conn)

		srv := server.Get("rpc-server-1")
		testdata.RegisterGreeterServer(srv, &GreeterServerImpl{})

		g := rgin.Gin()
		g.GET("/hello", func(ctx *gin.Context) {
			defer func() {
				if r := recover(); r != nil {
					s := GetTraceStack()
					fmt.Println(r)
					fmt.Println(s)
				}
			}()
			c := ctx.Request.Context()

			// http client
			hc := rhttplib.GetHttpClient("ngo.client.http.default")
			_, err := hc.Get("https://www.163.com").Do(c)
			if err != nil {
				fmt.Printf("get error:%s", err)
			}

			// db client
			dbClient := rgorm.GetDBClient("ngo.client.gorm.mysql")
			dbClient.WithContext(c).Exec("update user set age=22 where id=1")

			// redis client
			r := rredis.Get("ngo.client.redis.default")
			r.Get(c, "111")

			// kafka producer
			producer := rkafka.GetProducer("ngo.client.kafka.default")
			producer.SyncSend(c, "ngotest", "1234")

			// rpc client
			greeterClient.SayHello(c, &testdata.HelloRequest{Name: "world"})

			ctx.String(http.StatusOK, "hello world!")
		})
		return nil
	})
	app.Startup()

}

func GetTraceStack() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
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
