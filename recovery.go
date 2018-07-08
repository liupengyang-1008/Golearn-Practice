package main

import (
	"fmt"
)

func main() {
	var a int = 888

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {

			a = 999
			fmt.Println(a)

		}
	}()

	if a > 1000 {
		fmt.Println(a)
		panic("a is to large")
	}
	fmt.Println(a)
	a = 10
	fmt.Println(a)

}
