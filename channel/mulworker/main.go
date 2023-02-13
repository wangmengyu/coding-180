package main

import (
	"fmt"
	"time"
)

func createWorker(i int) chan int {
	ch := make(chan int)
	go doWorker(ch, i)
	return ch
}

func doWorker(ch chan int, i int) {
	for n := range ch {
		fmt.Printf("receive %c from worker %d\n", n, i)
	}
}

// chanDemo
func chanDemo() {
	// 提供10个channel,
	// 每个通道  分别绑定一个处理接收到的数据的goroutine, 对接收到的数据做处理.(绑定goroutine) , 各个协程里的每个元素独立工作 没有前后处理要求.
	chList := make([]chan int, 0)
	for i := 0; i < 10; i++ {
		ch := createWorker(i)
		chList = append(chList, ch)
	}

	for i := 0; i < 10; i++ {
		chList[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		chList[i] <- 'A' + i
	}

	time.Sleep(1 * time.Second)

}

func main() {

	chanDemo()
}
