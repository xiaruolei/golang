package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int)  {
	for {
		time.Sleep(1 * time.Second)
		n, ok := <-c
		if ok {
			fmt.Printf("worker %d received %d\n", id, n)
		}
	}
}

func createWorker (id int) chan<- int{
	c := make(chan int)
	go worker(id, c)

	return c
}

func main() {
	//var c1, c2 = generator(), generator()
	//w := createWorker(0)
	//n := 0
	//
	//for {
	//	//n := 0
	//	select {
	//	case n = <- c1:
	//		fmt.Println("11111111111")
	//		w <- n
	//	case n = <- c2:
	//		fmt.Println("22222222222")
	//		w <- n
	//	case w <- n:
	//	}
	//}

	//var c1, c2 = generator(), generator()
	//var worker = createWorker(0)
	//
	//n := 0
	//hasValue := false
	//for {
	//	var activeWorker chan<- int
	//	if hasValue {
	//		activeWorker = worker
	//	}
	//	select {
	//	case n = <- c1:
	//		hasValue = true
	//	case n = <- c2:
	//		hasValue = true
	//	case activeWorker <- n:
	//		hasValue = false
	//	}
	//}

	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	tick := time.Tick(time.Second)
	var values []int
	tm := time.After(10 * time.Second)

	n := 0
	for {
		var activeWorker chan<- int
		var activeValue int

		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <- c1:
			values = append(values, n)
		case n = <- c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <- time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len =", len(values))
		case <- tm:
			fmt.Println("Bye Bye")
			return
		}
	}

}

