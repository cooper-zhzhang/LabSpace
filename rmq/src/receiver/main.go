package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

var (
	Topic      = "my_topic"
	Grop       = "my_group_2"
	NameSpace  = "my_namespace"
	NameServer = "127.0.0.1:9876"
)

func receiveMessage(name string) rocketmq.PushConsumer {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{NameServer}),
		consumer.WithGroupName(Grop),
		consumer.WithNamespace(NameSpace), // 如果使用命名空间
		 consumer.WithConsumerOrder(true),
		// consumer.WithConsumeConcurrentlyMaxSpan(100),
	)
	if err != nil {
		panic(err)
		return nil
	}
	err = c.Subscribe(Topic, consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range msgs {
				n := rand.Int63n(8) + 1
				time.Sleep(time.Second * time.Duration(n))
				fmt.Printf("Finish %d %s Queue:[%d]\n", time.Now().Local().Unix(), msg.Body, msg.Queue.QueueId)
			}

			return consumer.ConsumeSuccess, nil
		})
	if err != nil {
		fmt.Println("订阅失败: " + err.Error())
		return c
	}

	if err = c.Start(); err != nil {
		fmt.Println("启动消费者失败: " + err.Error())
		return c
	}
	fmt.Println("Consumer started")
	// defer c.Shutdown()
	return c
}

func main() {
	// 创建消费者
	rand.Seed(time.Now().UnixNano())
	pid := os.Getpid()
	c1 := receiveMessage(strconv.Itoa(pid))
	defer c1.Shutdown()
	// c2 := receiveMessage("2")
	// defer c2.Shutdown()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
