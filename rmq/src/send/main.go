package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var (
	Topic      = "my_topic"
	Grop       = "my_group"
	NameSpace  = "my_namespace"
	NameServer = "127.0.0.1:9876"
)

// type QueueSelector interface {
// 	Select(msg *primitive.Message, mqs []*primitive.MessageQueue, lastBrokerName string) *primitive.MessageQueue
// }

type QueueSelectorImp struct{}

func (QueueSelectorImp) Select(msg *primitive.Message, mqs []*primitive.MessageQueue, lastBrokerName string) *primitive.MessageQueue {
	id, _ := strconv.Atoi(msg.GetShardingKey())
	// fmt.Println("name:", lastBrokerName, len(mqs))
	return mqs[id%len(mqs)]
}

func main() {
	// 创建一个生产者客户端
	p, err := rocketmq.NewProducer(
		producer.WithRetry(2), // 发送消息失败时重试的次数
		producer.WithNameServer([]string{NameServer}),
		producer.WithGroupName(Grop),
		producer.WithNamespace(NameSpace),
		producer.WithQueueSelector(QueueSelectorImp{}),
	)
	if err != nil {
		panic(err)
	}

	// 启动生产者
	if err = p.Start(); err != nil {
		panic(err)
	}
	defer p.Shutdown()

	index := 0
	for {
		// 创建消息
		// pid := os.Getpid()
		index = index + 1
		key := index % 3
		content := fmt.Sprintf("key: [%d] Hello %d", key, index)
		msg := primitive.NewMessage(
			Topic,
			[]byte(content),
		)
		msg = msg.WithShardingKey(strconv.Itoa(key))

		// 同步发送消息
		res, err := p.SendSync(context.Background(), msg)
		if err != nil {
			fmt.Printf("Send message error: %s\n", err)
		} else {
			fmt.Printf("Send message success, result: %s\n", res.String())
		}
		if index%10 == 0 {
			time.Sleep(time.Second)
		}
	}
}
