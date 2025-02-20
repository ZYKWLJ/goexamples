package main

import (
	"errors"
	"fmt"
)

func testError(arg int) (int, error) {
	if arg < 0 {
		return -1, errors.New("arg error,can't less than zero")
	} else {
		fmt.Println("arg greater than zero is valid :", arg)
		return arg, nil
	}

}
func Error(arg int) string {
	if arg > 10 {
		return fmt.Sprintf("arg greater than 10 is error\n")
	} else {
		return fmt.Sprintf("arg less  than 10 is valid\n")
	}
}

func main() {
	fmt.Println("这里验证go的错误机制处理")
	fmt.Println("使用内置的go的错误机制")
	testError(3)
	var a, b = testError(-1)
	fmt.Println(a, b)
	fmt.Println("使用自定义的错误机制")
	var c, d = Error(100), Error(9)
	fmt.Println(c, d)
}
