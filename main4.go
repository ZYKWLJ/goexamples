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

// 测试接口专用的文件
func main() {
	r := newr(4, 3)
	c := newc(4)
	fmt.Println(two(r))
	fmt.Println(two(c))
}

func two(s geometry) (float64, float64) {
	return s.area(), s.perimeter()
}
