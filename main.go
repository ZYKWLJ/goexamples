package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	a := "ssss"
	println(a)
	var i int = 1
	for {
		fmt.Printf("i=%d\n", i)
		if i == 5 {
			break
		}
		i++
	}
	if c := 5; i == 5 {
		fmt.Println(c)
		c++
		c--
		fmt.Println(c)
	}
	//switch用来判断类型
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	fmt.Println("数组测试")
	//var arr = make([]int, 20)
	var arr [10]int
	for i := 0; i < 10; i++ {
		//arr = append(arr, i)
		arr[i] = i
	}
	fmt.Println(arr)
	var d [5]int
	b := []int{1, 2, 3, 4, 5}
	var twoD [2][3]int
	fmt.Println(twoD, b, d)
	e := []int{2, 3, 4, 5, 7}
	copy(b, e)
	fmt.Println(b)
	b = append(b, 3)
	fmt.Println(b)

	fmt.Println("映射测试")
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(m)
	f := "b"
	if ans, ok := m[f]; ok {
		fmt.Println("存在键:", f, "\t其映射值为:", ans, "ok:", ok)
	} else {
		fmt.Println("不存在键", f, ans, "ok:", ok)
	}

	fmt.Println("实现的关联容器，类似于C++中的四类")
	fmt.Println("实现的set")
	set := make(map[int]struct{})
	set[1] = struct{}{}
	set[1] = struct{}{}
	set[1] = struct{}{}
	set[2] = struct{}{}
	fmt.Println(set)
	fmt.Println("实现的multiset")
	multiset := make(map[int]int)
	multiset[1] = 1
	multiset[1]++
	fmt.Println(multiset)
	fmt.Println("实现的multimap")
	multimap := make(map[int][]int) //一个键对应一个值的集合即可！
	multimap[1] = []int{1, 2, 3}
	multimap[1] = append(multimap[1], 0)
	fmt.Println(multimap)
	fmt.Println("迭代器ranges的使用")
	arr1 := []int{1, 2, 3}
	fmt.Println("迭代数组、切片")
	for i, v := range arr1 {
		fmt.Println(i, v)
	}
	fmt.Println("迭代映射")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("迭代切片")
	for i, char := range "123456789abcdefghijklmnoqprstuvwxyz" {
		fmt.Println(i, char, string(char))
		//rune2string := string(char)
		//fmt.Println(i, rune2string)
	}
	for i, c := range "go" {
		fmt.Println(i, string(c))
	}
}
