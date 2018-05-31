package main

import (
	"fmt"
)

type treeNode struct {
	Value       int
	Left, Right *treeNode
}

func main() {
	var root treeNode
	root = treeNode{Value: 3}
	root.Left = &treeNode{}
	root.Right = &treeNode{5, nil, nil}
	root.Right.Left = new(treeNode)
	root.Left.Right = creatNode(2)
	nodes := []treeNode{
		{Value: 2},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.print()
	root.Right.Left.setValue(4)
	root.Right.Left.print()

	var pNode *treeNode
	pNode.setValue(222)

	fmt.Println("----------------")
	root.traverse()
}

//局部变量可能在堆上 也可能在栈上
func creatNode(value int) *treeNode {
	return &treeNode{Value: value}
}

func (node *treeNode) print() {
	fmt.Println(node.Value)
}

func (node *treeNode) setValue(val int) {
	if node == nil {
		fmt.Println("Setting value to nil node .Ingored")
		return
	}
	node.Value = val
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.Left.traverse()
	node.print()
	node.Right.traverse()
}
