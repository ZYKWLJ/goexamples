package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		c1 <- "received meg from goroutine 1"
	}()

	now := time.Now()
	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(time.Second * 2):
		fmt.Println("msg from goroutine 1 timeout...", time.Since(now), "s")
	}

	now1 := time.Now()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "received meg from goroutine 2"
	}()
	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(time.Second * 1):
		fmt.Println("msg from goroutine 2 timeout...", time.Since(now1), "s")
	}
}
