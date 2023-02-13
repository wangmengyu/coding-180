package main

import (
	"fmt"
	"time"
)

func main() {
	//新建管道. 往里面塞数据,. 不消耗, 造成死锁
	ch := make(chan struct{})

	go consume(ch)
	ch <- struct{}{}
	time.Sleep(1 * time.Second)
	ch <- struct{}{}

	fmt.Println("done")
	time.Sleep(1 * time.Second)

}

func consume(ch chan struct{}) {
	for n := range ch { // 本身就是一种阻塞式的等待通道内的新元素
		fmt.Println(n)
	}
}
