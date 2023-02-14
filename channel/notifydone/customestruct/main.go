package main

import "fmt"

type worker struct {
	in   chan int
	done chan struct{}
}

func createWorker(i int) worker {
	w := worker{
		in:   make(chan int),      // 初始化管道
		done: make(chan struct{}), // 初始化完成信号管道
	}
	go doWorker(&w, i)
	return w
}

func doWorker(w *worker, i int) {
	for n := range w.in {
		fmt.Printf("receive %c from worker %d\n", n, i)
		// 为什么马上传回去?
		//  不马上传回就变成死循环了.
		w.done <- struct{}{}
	}

}
func main() {
	// 建立10个管道, 并且绑定各自的协程
	// 每个管道输入一个字符. 并行的进行字符的输出.
	// 在管道内完成任务告知外部协程完成信号,
	// 等到所有协程都发送了完成信号结束运行

	workerList := make([]worker, 0)
	for i := 0; i < 10; i++ {
		workerList = append(workerList, createWorker(i))
	}
	for i := 0; i < 10; i++ {
		// 对每个管道输入一个字符
		workerList[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		<-workerList[i].done
	}

	fmt.Println("done")

}
