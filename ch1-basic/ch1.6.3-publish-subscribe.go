package main

import (
	"fmt"
	"strings"
	"time"

	"chai2010.cn/gobook/ch1-basic/pubsub"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello,  world!")
	p.Publish("hello, golang!")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}

// output
// all: hello,  world!
// all: hello, golang!
// golang: hello, golang!