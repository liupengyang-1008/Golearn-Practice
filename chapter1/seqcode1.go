package main

import "fmt"

func main() {
	constdata()
	declare()

}

func constdata() {
	fmt.Println("func constdata start!")
	const Pi float64 = 3.1415926
	const zero = 0.0
	const (
		size int64 = 1024
		eof        = -1
	)
	const u, v float32 = 0, 3
	const a, b, c = 3, 4, "lod"
	const mask = 1 << 3
	const (
		c0 = iota
		c1 = iota
	)

	fmt.Println(Pi)
	fmt.Println(zero)
	fmt.Println(size)
	fmt.Println(eof)
	fmt.Println(u, v, a, b, c, c0, c1, mask)
	fmt.Println("func constdata end!")
}

func declare() {

	fmt.Println("func declare start!")
	var v1 int = 10
	v2 := 30

	fmt.Println(v1, v2)
	v1, v2 = v2, v1
	fmt.Println(v1, v2)
	//匿名变量
	_, _, nike := getname()
	fmt.Println(nike)

	fmt.Println("func declare end!")
}

func getname() (fname string, sname string, tname string) {
	return "may", "Chan", "Child lod"
}
