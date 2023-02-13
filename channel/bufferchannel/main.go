package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个带buffer的channel。 他最多能承受无消耗的N个元素
	// 在管道内输入a开始的10个字母, 并且用worker来进行消耗打印出来
	ch := make(chan int, 3)
	go consume(ch)
	for i := 0; i < 10; i++ {
		ch <- 'a' + i
	}
	time.Sleep(1 * time.Second)
}

func consume(ch chan int) {
	for n := range ch {
		fmt.Printf("%c\n", n)
	}
}
