package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(chan int)
	go func() {
		fmt.Println("开始通道传输...")
		time.Sleep(3 * time.Second)
		m <- 1
		fmt.Println("完成通道传输...")

		fmt.Println("通道信息接受在主线程中完成")

	}()

	fmt.Println("开始通道接受...")
	s := <-m
	fmt.Println(s)
	fmt.Println("完成通道接受...")

	str := fmt.Sprintf("这是主线程的消息,需要等待其他协程将信息传给我了，我才能继续执行")
	fmt.Printf("\033[31m %s \033[0m", str)

}
