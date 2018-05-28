// BingFa
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	thread()
	channel()

}
func count2(ch chan int) {
	fmt.Println("Counting ", ch)
	ch <- 1
}
func channel() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go count2(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}

var counter int = 0

func Count1(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()

}

func thread() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count1(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}
