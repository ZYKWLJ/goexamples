package main

import (
	"fmt"
	"math"
)

type listNode struct {
	next *listNode
	val  float64
}

func newList(val float64) *listNode {
	return &listNode{
		next: nil,
		val:  val,
	}
}

func main() {
	fmt.Println("演示泛型的使用", math.Pi)
	fmt.Println("下面是一个链表")
	head := newList(math.Pi)
	head.next = newList(math.Pi * math.Pi)
	head.next.next = newList(math.Pi * math.Pi * math.Pi)
	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
}
