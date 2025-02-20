package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("演示一个协程同时等待多个协程~")
	a, b := make(chan string), make(chan string)
	go func() {
		start := time.Now()
		time.Sleep(time.Second * 1)
		str := fmt.Sprintf("这是第1个协程写来的数据，用时 %fms", float64(time.Since(start).Milliseconds()))
		a <- str
	}()

	go func() {
		start := time.Now()
		time.Sleep(time.Second * 2)
		str := fmt.Sprintf("这是第2个协程写来的数据，用时 %fms", float64(time.Since(start).Milliseconds()))
		b <- str
	}()
	for i := 0; i < 2; i++ {
		select {
		case str := <-a:
			{
				fmt.Println(str)
			}
		case str := <-b:
			fmt.Println(str)

		}
	}
}
