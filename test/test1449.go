package main

import "fmt"

func main() {
	ints := make([]int, 100)
	for i := range ints {

		fmt.Println(i)
	}
}
