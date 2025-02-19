package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("1,4的和\\差")
	ret1, ret2 := functiontest1(1, 4)
	fmt.Println(ret1, ret2)
	functiontest2(1, 2, 3, 4)
	fmt.Println("测试通过切片传参")
	slices := []int{100, 200, 300, 400}
	functiontest2(slices...)
	fmt.Println("测试返回函数的函数(闭包，匿名函数)")
	nextInt := retFunc()
	fmt.Println(nextInt(2))
	fmt.Println(nextInt(3))
	fmt.Println(nextInt(4))
	fmt.Println(retFunc()(4))
	fmt.Println("测试闭包,必须在前面定义闭包")
	var fib func(a int) int
	fib = func(a int) int {
		if a == 0 || a == 1 {
			return 1
		}
		return fib(a-1) + fib(a-2)
	}
	fmt.Println(fib(20))
	fmt.Println(fib1(5))

	fmt.Println("演示go中的rune")
	const s = "你好呀"
	fmt.Println(len(s)) //长度为9字节
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	//数量为3个，表示字符数，而非字节数，因为一个字符咱占用一个或多个字节
	fmt.Println("\nRune count:", utf8.RuneCountInString(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		//examineRune(runeValue)
	}
	t := "你"
	fmt.Println(len(t))
	var str string = "你好"
	fmt.Println("结构体知识")
	fmt.Println(people{"郑阳康", 20})
	fmt.Println(&people{"郑阳康", 20})
	fmt.Println(people{age: 20})
	fmt.Println(&str)
}

type people struct {
	name string
	age  int
}

func functiontest1(a, b int) (int, int) {
	fmt.Println("测试参数压缩与多返回值")
	return a + b, a - b
}

func functiontest2(nums ...int) {
	for i, n := range nums {
		fmt.Println(i, n)
	}
}

func retFunc() func(a int) int {
	i := 1
	return func(a int) int {
		i++
		return i + a
	}
}

func fib1(a int) int {
	if a < 2 {
		return 1
	}
	ret := fib1(a-1) + fib1(a-2)
	return ret
}
