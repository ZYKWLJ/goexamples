package main

import (
	"fmt"
	"time"
)

// 演示通过select+default来完成非阻塞操作，本质是select如果没有接收到消息直接就退出了
func main() {
	c := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c <- "ping"
	}()

	fmt.Println("test sent messages...")
	select {
	case message := <-c:
		println("sent msg: ", message)
	default:
		//加上这个分支后，这个选择器不会一直阻塞等待通道里面的消息，而是会直接打印退出，但是不加就会一直等待
		println("haven't sent msg,default...")
	}
}
