package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	// 新建channel推送数据
	ch := make(chan int)
	// 消耗数据,一定要在送数据之前使用。否则还是会死锁
	go consume(ch)
	ch <- 1
	ch <- 2
	time.Sleep(1 * time.Second)
}

func consume(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
func main() {
	chanDemo()
}
