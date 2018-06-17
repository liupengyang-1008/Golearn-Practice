//使用channel 等待gorounting结束
package main

import (
	"fmt"
	"sync"
)

func main() {
chanDemo()
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i:=0 ;i<10 ;i++ {
		workers[i] = createWorker(i,&wg)
	}

	for i,worker := range workers {
		worker.in <- 'a' +i
		//<-workers[i].done
		wg.Add(1)
	}

	for i,worker := range workers {
		worker.in <- 'A' +i
		//<-workers[i].done
		wg.Add(1)
	}

    wg.Wait()

////	wait all done
//	for _,worker := range workers {
//		<-worker.done
//	}
}

func createWorker(id int ,wg *sync.WaitGroup) worker {
	w := worker{
		in:make(chan int),
		wg:wg,
	}
	go doworker(id,w.in,w.wg )
	return  w
}

func doworker(id int, c chan int,wg *sync.WaitGroup ) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id,n)
		//
		// go func() {	done <- true}()
		wg.Done()

	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup

} 