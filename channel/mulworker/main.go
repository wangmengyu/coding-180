package main

import (
	"fmt"
	"time"
)

// doWork 一直从管道收数据进行处理
func doWork(i int, ch chan int) {
	for v := range ch {
		fmt.Printf("receive %c from worker %d\n", v, i)
	}
}

// createWorker  创建worker
func createWorker(i int) chan int {
	ch := make(chan int)
	go doWork(i, ch)
	return ch
}

// chanDemo
func chanDemo() {
	// 建立10个通道的数组
	var chList [10]chan int

	//每个通道都可以消耗通道里的元素
	for i := 0; i < 10; i++ {
		ch := createWorker(i)
		chList[i] = ch

	}
	// 向通道内按照顺序放10个a开头的字符
	for i := 0; i < 10; i++ {
		chList[i] <- 'a' + i
	}

	//  向通道内按照顺序放10个A开头的字符
	for i := 0; i < 10; i++ {
		chList[i] <- 'A' + i
	}

	time.Sleep(1 * time.Second)

}
func main() {

	chanDemo()
}
