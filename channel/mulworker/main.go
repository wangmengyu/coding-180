package main

import (
	"fmt"
	"time"
)

func createWorker(i int) chan int {
	ch := make(chan int, 0)
	go doWorker(ch, i)
	return ch

}

func doWorker(ch chan int, i int) {
	for n := range ch {
		fmt.Printf("read %c from worker %d\n", n, i)
	}
}

// chanDemo
func chanDemo() {
	// 提供10个channel,
	// 每个通道分别处理接收到的数据.(绑定goroutine)

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
