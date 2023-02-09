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
	"strconv"
	"time"

	"github.com/NetEase-Media/easy-ngo/application"
	_ "github.com/NetEase-Media/easy-ngo/application/r/rconfig"
	"github.com/NetEase-Media/easy-ngo/application/r/rkafka"
	"github.com/NetEase-Media/easy-ngo/clients/xkafka"
)

const topic = "test"

var (
	consumer    *xkafka.Consumer
	producer    *xkafka.Producer
	messageChan = make(chan xkafka.ConsumerMessage, 1000)
)

// go run . -c ./app.yaml
func main() {
	app := application.Default()
	app.Initialize()
	app.Startup()

	consumer = rkafka.GetConsumer("ngo.client.kafka.demo01")
	//consumer.AddListener(topic, &listener{})
	consumer.AddBatchListener(topic, &batchListener{})
	consumer.Start()

	producer = rkafka.GetProducer("ngo.client.kafka.demo01")
	for i := 0; i < 98; i++ {
		producer.Send(topic, "hello world!"+strconv.Itoa(i), nil)
	}
	go func() {
		for {
			r := <-messageChan
			fmt.Print(r.Value, " ", r.Partition, " ", r.Offset)
		}
	}()
	time.Sleep(10 * time.Second)
}

type listener struct {
	xkafka.Listener
}

func (l *listener) Listen(message xkafka.ConsumerMessage, ack *xkafka.Acknowledgment) {
	messageChan <- message
	ack.Acknowledge()
}

type batchListener struct {
	xkafka.BatchListener
}

func (l *batchListener) Listen(messages []xkafka.ConsumerMessage, ack *xkafka.Acknowledgment) {
	fmt.Print("batch size is ", len(messages))
	for _, msg := range messages {
		messageChan <- msg
	}
	ack.Acknowledge()
}

func (l *batchListener) BatchCount() int {
	return 10
}
