package main

import "fmt"

func main() {

	ch := make(chan int, 1) // 创建一个容量为 1 的有缓冲通道,就算没有其他的协程接收也不会出错，因为自己有缓冲，可以不必等待，自己就可以接受
	//ch := make(chan int//不加的话，会导致一直等待其他协程来接收，会阻塞出错！
	// 向通道发送数据，由于缓冲区有空间，不会阻塞
	ch <- 1
	// 从通道接收数据
	num := <-ch
	fmt.Println("Received:", num)
}
