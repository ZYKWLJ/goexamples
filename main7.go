package main

import (
	"fmt"
	"time"
)

func testGoRountine1(m int) {
	for m > 0 {
		fmt.Printf("\033[31m %d \033[0m", m)
		m--
	}
}
func testGoRountine2(n int) {
	for n > 0 {
		fmt.Printf("\033[32m %d \033[0m", n)
		n--
	}
}
func main() {
	fmt.Println("这是阻塞执行的")
	testGoRountine1(1000)
	testGoRountine2(2000)
	fmt.Println("\n\n这是并发执行的")
	go testGoRountine1(1000)
	go testGoRountine2(2000)
	time.Sleep(1000 * time.Millisecond)
}
