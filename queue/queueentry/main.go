package main

import (
	"Golearn-Practice/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q, &q, q.Pop())
	fmt.Println(q, &q, q.Pop())
	fmt.Println(q, &q, q.IsEmpty())
	fmt.Println(q, &q, q.Pop())
	fmt.Println(q, &q, q.IsEmpty())
}
