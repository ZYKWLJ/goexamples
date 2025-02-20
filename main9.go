package main

import (
	"fmt"
	"time"
)

type SYNACK1 struct {
	SYN string
	ACK string
}

// 一个小小的实验演示通道发送和接受是阻塞的
func waitMeReceive(mesg chan SYNACK1) {
	// 从通道中接收消息
	synack := <-mesg
	fmt.Println(synack.SYN)
	time.Sleep(1 * time.Second)
	fmt.Println(synack.ACK)
	mesg <- SYNACK1{SYN: "1", ACK: "0"}
}

func main() {
	// 创建一个无缓冲的通道
	mesgqueue := make(chan SYNACK1)
	// 启动一个 goroutine 来处理通道中的消息
	go waitMeReceive(mesgqueue)
	synack := SYNACK1{SYN: "你好呀，我发来的电报你收到了吗？", ACK: "是的，我收到了！"}
	// 向通道中发送消息
	mesgqueue <- synack
	<-mesgqueue //这里的通道接受操作是必不可少的，因为这是同步阻塞进行的，不然主协程会直接结束了！
}
