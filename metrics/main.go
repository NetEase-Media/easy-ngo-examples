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
	"time"

	_ "github.com/NetEase-Media/easy-ngo/application/r/rconfig"

	_ "github.com/NetEase-Media/easy-ngo/application/r/rprometheus"

	"github.com/NetEase-Media/easy-ngo/application"

	"github.com/NetEase-Media/easy-ngo/application/r/rgin"
	"github.com/NetEase-Media/easy-ngo/application/r/rgorm"
	"github.com/NetEase-Media/easy-ngo/application/r/rhttplib"
	"github.com/NetEase-Media/easy-ngo/application/r/rkafka"
	"github.com/NetEase-Media/easy-ngo/application/r/rmemcache"
	"github.com/NetEase-Media/easy-ngo/application/r/rms/grpc/client"
	"github.com/NetEase-Media/easy-ngo/application/r/rms/grpc/server"
	"github.com/NetEase-Media/easy-ngo/application/r/rredis"
	"github.com/NetEase-Media/easy-ngo/clients/xkafka"
	"github.com/NetEase-Media/easy-ngo/microservices/testdata"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// go run . -c ./app.yaml
func main() {
	//启动metrics
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe("0.0.0.0:19933", nil)

	app := &application.Application{}
	// app初始化完成回调
	app.Initialize(func() error {

		conn := client.Get("rpc-client-1")
		greeterClient := testdata.NewGreeterClient(conn)

		srv := server.Get("rpc-server-1")
		testdata.RegisterGreeterServer(srv, &GreeterServerImpl{})

		consumer := rkafka.GetConsumer("ngo.client.kafka.default")
		consumer.AddListener("ngotest", &Listener{})
		go consumer.Start()

		g := rgin.Gin()
		g.GET("/hello", func(ctx *gin.Context) {
			defer func() {
				if r := recover(); r != nil {
					s := GetStack()
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
			db := rgorm.GetDBClient("ngo.client.gorm.mysql")
			db.WithContext(c).Exec("update user set age = ? where id= ?", 22, 1)
			var t User
			db.WithContext(c).Raw("select * from user where id = ?", 1).Find(&t)

			// redis client
			r := rredis.Get("ngo.client.redis.default")
			r.Set(c, "key1", "value1", time.Second*5)
			r.Get(c, "key1")

			// kafka producer
			p := rkafka.GetProducer("ngo.client.kafka.default")
			p.SyncSend(c, "ngotest", "1234")

			// memcached
			m := rmemcache.GetClient("ngo.client.memcache.default")
			m.Set(c, "key1", "value1")
			m.Get(c, "key1")

			// rpc client
			greeterClient.SayHello(c, &testdata.HelloRequest{Name: "world"})

			ctx.String(http.StatusOK, "hello world!")
		})
		return nil
	})
	app.Startup()

}

func GetStack() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var _ xkafka.Listener = (*Listener)(nil)

type Listener struct {
}

func (l Listener) Listen(msg xkafka.ConsumerMessage, _ *xkafka.Acknowledgment) {
	fmt.Sprintf("topic %s, message: %s\n", msg.Topic, msg.Value)
}

type GreeterServerImpl struct {
	testdata.UnimplementedGreeterServer
}

func (g GreeterServerImpl) SayHello(ctx context.Context, request *testdata.HelloRequest) (*testdata.HelloReply, error) {
	fmt.Sprintf("hello %s \n", request.Name)
	return &testdata.HelloReply{
		Message: fmt.Sprintf("hello %s", request.Name),
	}, nil
}
