package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int //  存放业务数据的通道
	done func()   // 执行完毕后通知wg完成
}

func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(i, w)
	return w

}

func doWorker(i int, w worker) {
	for n := range w.in {
		fmt.Printf("receive %c from %d worker\n", n, i)
		w.done()
	}
}

func main() {

	// 10个通道发送
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	// 传输数据前先设置号处理总数
	wg.Add(10)
	for i, w := range workers {
		w.in <- 'a' + i
	}
	//  等待执行完成
	wg.Wait()

}
