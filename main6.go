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

type argError struct {
	arg int
}

func (arg *argError) Error() string {
	if arg.arg > 10 {
		return fmt.Sprintf("%d, arg greater than 10 is error\n", arg.arg)
	} else {
		return fmt.Sprintf("%d, arg less  than 10 is valid\n", arg.arg)
	}
}

func main() {
	fmt.Println("这里验证go的错误机制处理")
	fmt.Println("使用内置的go的错误机制")
	testError(3)
	var a, b = testError(-1)
	fmt.Println(a, b)
	fmt.Println("使用自定义的错误机制")
	var c, d = argError{100}, argError{9}
	fmt.Println(c.Error(), d.Error())
}
