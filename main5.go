package main

import (
	"fmt"
	"math"
)

// 链表结构体定义
type listNode[K comparable] struct {
	next *listNode[K]
	val  K
}

// 初始化链表
func newList[K comparable](val K) *listNode[K] {
	return &listNode[K]{
		next: nil,
		val:  val,
	}
}

func main() {
	fmt.Println("演示泛型的使用", math.Pi)
	fmt.Println("下面是一个int类型链表")
	head := newList(math.Pi)
	head.next = newList(math.Pi * math.Pi)
	head.next.next = newList(math.Pi * math.Pi * math.Pi)
	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
	fmt.Println("下面是一个string类型链表")
	head1 := newList("你好呀")
	head1.next = newList("我是")
	head1.next.next = newList("郑阳康")
	for head1 != nil {
		fmt.Println(head1.val)
		head1 = head1.next
	}
}
