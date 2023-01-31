package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generator 源源不断， 有间隔的往通道内发送数据。
func generator() chan int {
	// 通道初始化
	ch := make(chan int)

	// 源源不断的向通道内发送递增的整数
	go func() {
		i := 0
		for {
			// 每隔1500随机秒毫秒数等待
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			ch <- i
			i++
		}
	}()
	return ch

}

// 创建一个接受数据的统一通道来处理数据 . 每个元素并发的处理
func createWorker() chan int {
	ch := make(chan int)
	go doWorker(ch)
	return ch

}

// 接受来自c1,c2 的数据进行进一步处理
func doWorker(ch chan int) {
	for n := range ch {
		// 消耗速度放慢
		time.Sleep(1 * time.Second)
		fmt.Println("Received ", n)
	}
}

func main() {
	//  有两个通道， 同时从两个通道获取数据。 哪个先来哪个先处理
	c1 := generator()
	c2 := generator()

	// 需求: 两个通道内的数据全部收集到一个worker管道内并发的将每个元素处理掉
	n := 0                    // 初始化处理的变量
	numList := make([]int, 0) // 代处理数据队列.
	worker := createWorker()  // 初始化好的worker, 等待有新值的时候被使用
	//10秒后会发送消息到通道
	tm := time.After(10 * time.Second)
	// 每秒种定时查看队列长度定时器
	ti := time.Tick(1 * time.Second)

	for {
		var activeWorker chan int
		var actVal int

		timeoutTm := time.After(800 * time.Millisecond)
		if len(numList) > 0 {
			// 待处理队列有记录需要处理
			activeWorker = worker
			actVal = numList[0] // get first elem
		}

		select {
		case n = <-c1:
			numList = append(numList, n)
		case n = <-c2:
			numList = append(numList, n)
		case activeWorker <- actVal:
			numList = numList[1:] //处理掉,移除队列
		case <-timeoutTm: // 每次循环800 毫秒没有输出就打印timeout
			fmt.Println("timeout")
		case <-ti: // 每秒查看一下队列长度.
			fmt.Println("len of queue:", len(numList))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}

}
