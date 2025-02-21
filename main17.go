package main

import "fmt"

func main() {
	//	关闭通道的小实验
	var work, block = make(chan int, 5), make(chan bool)
	go func() {
		// 使用 for...range 循环从通道接收数据
		for meg := range work {
			fmt.Println("Success! received msg ", meg)
		}
		fmt.Println("received all msg! Work closed.")
		block <- false
	}()
	for i := 0; i < 3; i++ {
		work <- i
		fmt.Println("sent msg", i)
	}
	close(work)
	fmt.Println("sent over!")

	<-block //一直阻塞主线程的作用！
}
