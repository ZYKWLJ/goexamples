package main

import (
	//"bufio"
	"fmt"
	"math"
)

type rect struct {
	w int
	h int
}
type circle struct {
	radius int
}
type geometry interface {
	area() float64
	perimeter() float64
}

func newr(w, h int) rect {
	return rect{w: w, h: h}
}
func newc(r int) circle {
	return circle{radius: r}
}
func (r rect) area() float64 {
	return float64(r.h * r.w)
}
func (r rect) perimeter() float64 {
	return 2 * float64(r.w+r.h)
}
func (c circle) area() float64 {
	return math.Pi * float64(c.radius*c.radius)
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * float64(c.radius*c.radius)
}

// 接口和结构体嵌入成为更复杂的接口
type nameAndShape struct {
	geometry
	name any
}

// 测试接口专用的文件
func main() {
	r := newr(4, 3)
	c := newc(4)
	r1 := &r
	fmt.Println(two(r1))
	fmt.Println(two(c))
	fmt.Println("结构体和接口嵌入，看起来就像没有名字的字段一样")
	NASr := nameAndShape{r, "这是矩形"}
	NASc := nameAndShape{c, "这是圆形"}
	fmt.Println(NASr, NASc, NASc.name, NASr.geometry.area())
}

// 注意接口类型尽量少作为作为指针传参
func two(s geometry) (float64, float64) {
	return s.area(), s.perimeter()
}
