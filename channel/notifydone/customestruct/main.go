package main

import "fmt"

type worker struct {
	in   chan int      // 处理数据的管道
	done chan struct{} // 通知完成一个元素的处理
}

func createWorker(i int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan struct{}),
	}
	go doWorker(i, w)
	return w

}

// doWorker 消耗管道内元素
func doWorker(i int, w worker) {
	for n := range w.in {
		fmt.Printf("receive %c from worker %d\n", n, i)
		w.done <- struct{}{} // 发送元素处理完毕标记, 做完一个发送一个处理完的信号。
	}

}
func main() {

	// 建立10个管道。
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	// 先输入小写字母 到每个管道
	for i, w := range workers {
		w.in <- i + 'a'
	}
	// 遍历done的管道接收掉处理完毕的信号
	for _, w := range workers {
		<-w.done
	}

}
