package main

import "fmt"

func fibonacci(size int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < size; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch) // 在函数内部关闭掉channel

}
func main() {
	// 建立一个10个缓冲区的CHANNEL
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for m := range ch {
		fmt.Println(m)
	}

}
