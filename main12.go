package main

import "fmt"

// 只写通道(只允许数据写进来)
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 只读通道(允许读取通道中的数据)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	fmt.Println("下面将演示数据进入只写通道函数，只读通道函数")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
