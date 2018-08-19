package main

import (
	"fmt"
	"time"
	"math/rand"
)


func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(  time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func generator() chan int {
	out :=make(chan int )
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out  <- i
			i ++
		}
	}()
	return out
}


func main() {
	var c1,c2 = generator(),generator()
	worker :=createWorker(0)
	n := 0
	haseValue := false
	var values
	for {
	var activeWorker chan<- int
	if haseValue {
		activeWorker = worker
	}
	select {
	case  n =  <- c1 :
		haseValue = true
		//fmt.Println("received from c1 :",n)
	case  n =  <- c2 :
		//fmt.Println("received from c2 :",n)
		haseValue = true
	case activeWorker  <- n :
		if haseValue {
			haseValue = false
		}

	}
	}

}

