package main

import (
	"fmt"
	"time"
)

func makeGen(done chan struct{}) chan string {
	ch := make(chan string)
	i := 0
	go func(i int) {
		for {
			// 源源不断生成字符串
			select {
			case <-done:
				fmt.Println("start clean task")
				time.Sleep(time.Duration(1000) * time.Millisecond)
				fmt.Println("clean up ")
				done <- struct{}{}
			case <-time.After(1 * time.Second): // 开始等待1秒
				fmt.Println("i=", i)
				time.Sleep(2 * time.Second)
				fmt.Println("after 2 second")
				ch <- fmt.Sprintf("generate i=%d", i)
				i++
			}
		}

	}(i)

	return ch
}

func main() {

	doneCh := make(chan struct{})
	ch := makeGen(doneCh)
	for i := 1; i <= 5; i++ {
		select {
		case v := <-ch:
			// get data from channel
			fmt.Println("v=", v)

		}
		i++
	}
	doneCh <- struct{}{} // 5次之后终止请求
	<-doneCh
	fmt.Println("done")

}
