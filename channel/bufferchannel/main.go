package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	for n := range ch {
		fmt.Printf("%c\n", n)
	}
}
func main() {
	// 定义一个带buffer的channel。 他最多能承受无消耗的N个元素
	ch := make(chan int, 3)
	go worker(ch)
	for i := 0; i < 10; i++ {
		ch <- i + 'a'
	}
	time.Sleep(1 * time.Second)
}
