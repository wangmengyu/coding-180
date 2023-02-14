package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
}

func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{in: make(chan int)}
	go doWorker(w, i, wg)
	return w
}

func doWorker(w worker, i int, wg *sync.WaitGroup) {
	for n := range w.in {
		fmt.Printf("receive %c from worker %d\n", n, i)
		wg.Done()
	}
}
func main() {

	// 10个通道发送字符
	// 并行的完成字符的处理
	var wg sync.WaitGroup
	workers := make([]worker, 0)

	for i := 0; i < 10; i++ {
		workers = append(workers, createWorker(i, &wg))
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	wg.Wait()

}
