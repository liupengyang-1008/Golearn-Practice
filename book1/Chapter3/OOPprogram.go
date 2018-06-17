// OOPprogram.go
package main

import (
	"fmt"
)

func main() {

	newlei()
	jiegouti()
}

//struct
func jiegouti() {
	rec1 := new(Rect)
	rec1.x = 1 ;rec1.y = 2 

}

type Rect struct {
	x, y          float64
	width, height float64
}

//rectangle   area
func (r *Rect) Area() float64 {
	return r.width * r.height
}

// 为类型新增方法
//define
type Integer int

// a < b = true
func (a Integer) Less(b Integer) bool {
	return a < b
}

//add 想修改值则需要传递指针
func (a *Integer) IntegerAdd(b Integer) {
	*a += b
}

//test
func newlei() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "is less than 2!")
	}
	a.IntegerAdd(4)
	fmt.Println("1 + 4 = ", a)

}
