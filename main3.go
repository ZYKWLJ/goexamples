package main

import "fmt"

// 测试方法专用
type area struct {
	w int
	h int
}

func newArea(w, h int) area {
	return area{
		w: w,
		h: h,
	}
}
func (a area) square() float64 {
	return float64(a.w * a.h)
}
func (a *area) changeSquare(w, h int) {
	a.w = w
	a.h = h
}
func main() {
	area := newArea(10, 20)
	fmt.Println(area)
	fmt.Println(area.square())
	area.changeSquare(10, 200)
	fmt.Println(area)
	fmt.Println(area.square())
}
