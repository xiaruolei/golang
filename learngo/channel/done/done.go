package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker)  {
	for {
		n, ok := <-w.in
		if ok {
			fmt.Printf("worker %d received %c\n", id, n)
			w.done()
		}

		//go func() {
		//	for {
		//		done <- true
		//	}
		//}()

	}

	//for n := range c {
	//	fmt.Printf("worker %d received %c\n", id, n)
	//	done <- true
	//}
}

type worker struct {
	in chan int
	//done chan bool
	//wg *sync.WaitGroup
	done func()

}

func createWorker (id int, wg *sync.WaitGroup) worker{
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWorker(id, w)
	return w
}

func chanDemo()  {
	//c := make(chan int)
	var workers [10]worker

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		//fmt.Println(<-workers[i].done)
	}

	//for _, worker := range workers {
	//	<-worker.wg
	//}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//fmt.Println(<-workers[i].done)
	}

	//for _, worker := range workers {
	//	<-worker.wg
	//}

	wg.Wait()


}


func main() {
	chanDemo()
}
