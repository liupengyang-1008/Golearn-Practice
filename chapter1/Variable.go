package main

import "fmt"

func main() {

	slice_go()
	slice2_go()
	map_go()
	Ctrl_flow()
	myfunc(1, 2, 3, 4)
	bibao()
}

//匿名函数
func bibao() {
	j := 5
	a := func() func() {
		var i int = 10
		i = i + 1
		j = j + 1
		return func() {
			fmt.Printf("i,j : %d ,%d \n", i, j)
		}

	}()
	a()
	j *= 2
	a()
	j *= 2
	a()
}

//Varargs 不定参数
func myfunc(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is a int data!")
		case string:
			fmt.Println(arg, "is a string data!")
		case int64:
			fmt.Println(arg, "is a int64 data!")
		default:
			fmt.Println(arg, "is a unknown data")

		}
	}
	fmt.Println("")
}

//flow control
func Ctrl_flow() {
	fmt.Println("-------------------------------------")

	if 1 > 0 {
		display(" 1 > 0 ")
	} else {
		display("1 <= 0")
	}

	//switch
	r := 4
	switch r {
	case 0:
		display("0")
	case 1:
		display("1")
	case 2:
		fallthrough
	case 3:
		display("3")
	case 4, 5, 6:
		display("4,5,6")
	default:
		display("Default")
	}

	//loop
	a := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
}

//map
func map_go() {
	fmt.Println("-------------------------------------")

	type personInfo struct {
		Id      string
		Name    string
		Address string
	}
	//declaration 声明
	var personDB map[string]personInfo
	// 创建 make,后面参数5表示存储能力为5
	personDB = make(map[string]personInfo, 5)

	//insert map data
	personDB["1"] = personInfo{"1", "bob", "road 1."}
	personDB["2"] = personInfo{"2", "jack", "road 12."}

	//search map data
	person, ok := personDB["2"]
	if ok {
		fmt.Println("found person ", person, "with id 2")

	} else {
		fmt.Println("Do not found person with id 2")
	}
	//delete map element
	delete(personDB, "1")

	fmt.Println(personDB["22"])

}

//切片元素
func slice2_go() {
	fmt.Println("-------------------------------------")
	mySlice2 := make([]int, 5, 10)
	mySlice2 = []int{2, 0, 3, 5, 6}
	fmt.Println("len of myslice2:", len(mySlice2))

	fmt.Println("cap of myslice2:", cap(mySlice2))

	mySlice2 = append(mySlice2, 1, 2, 3)
	fmt.Println("mySlice2:", mySlice2)
	mySlice := []int{2, 0, 3, 0, 6}
	// 从后面复制到前面
	copy(mySlice2, mySlice)
	fmt.Println("mySlice2:", mySlice2)

}

//切片初始化
func slice_go() {
	fmt.Println("-------------------------------------")
	//define a Array
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	//define a slice 前闭后开区间
	var mySlice []int = myArray[:5]

	fmt.Println("Element of myArray ：")
	for _, v := range myArray {
		fmt.Print(v, "")
	}

	fmt.Println("\nElement of mySlice ：")
	for _, v := range mySlice {
		fmt.Print(v, "")
	}
	fmt.Println("")
}

func display(s string) {
	fmt.Println(s)
}
