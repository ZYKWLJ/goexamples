package main

import "fmt"

func main() {
	fmt.Println("该实验演示效果")
	fmt.Println("1.一个非空的通道也是可以关闭的，而且关闭后，里面的值能够正常接受，只是不能再往里面发送值而已")
	fmt.Println("2.range的通道迭代法")
	queue := make(chan string, 20)
	queue <- "one"
	queue <- "two"
	close(queue)
	//queue <- "three"
	for elem := range queue {
		fmt.Println(elem)
	}
}
