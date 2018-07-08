// reflact
package main

import (
	"fmt"
	"reflect"
)

// 角色选择
type ChooseRole struct {
	RoleId   uint32 // 1 2 3
	RoleType string //tank
}

func main() {
	s := ChooseRole{
		RoleId:   1,
		RoleType: "tank",
	}
	test(&s)
}

func test(s interface{}) {
	msgType := reflect.TypeOf(s)
	msgElem := msgType.Elem()
	msgID := msgType.Elem().Name()
	fmt.Println("msgType:", msgType)
	fmt.Println("msgElem:", msgElem)
	fmt.Println("msgID:", msgID)
}
