package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	/* 这样写会输出无效字符。 当close之后。 通道内还在收很多字符
	for {
			v := <-ch
			fmt.Printf("%c\n", v)
		}
	*/
	//  正确应该用range过滤掉空字符
	for n := range ch {
		fmt.Printf("%c\n", n)
	}

}
func main() {
	// 定义一个channel
	ch := make(chan int, 3)
	defer close(ch)

	go worker(ch)
	for i := 0; i < 10; i++ {
		ch <- 'a' + i
	}

	time.Sleep(1 * time.Second)

}
