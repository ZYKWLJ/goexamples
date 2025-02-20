package main

import "fmt"

type SYNACK struct {
	SYN string
	ACK string
}

// 通道：语言级别的消息队列，是go语言具有极简的线程通信方式——基于消息队列通信方式
func main() {
	fmt.Println("创建一个通道")
	message := make(chan SYNACK, 10)
	fmt.Println("将信息传入通道里面")
	synack := SYNACK{SYN: "hey, this is my message, have you received it?", ACK: "Yes, I received it, and i will accept it ,thanks"}
	go func() { //这里是阻塞的发送和接受
		for i := 0; i < 10; i++ {
			message <- synack
		}
	}() //这是一个匿名函数

	ACK := <-message
	for i := 0; i < 10; i++ {
		fmt.Println(synack.SYN, " times= ", i)
		fmt.Println(ACK.ACK, " times= ", i)
	}
}
