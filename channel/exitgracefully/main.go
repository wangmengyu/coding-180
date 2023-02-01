package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeGen(serviceName string, done chan struct{}) chan string {
	ch := make(chan string)
	// 生成数据需要在独立的协程里完成

	go func() {
		i := 0
		for {
			select {
			case <-done: // 在执行的时候同时监听有没有done的信号， 有信号直接走退出流程。
				time.Sleep(2 * time.Second)
				fmt.Println("cleanup done")
				done <- struct{}{} // 完成后向done通知清理完成
				return
			case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
				ch <- fmt.Sprintf("service %s receive %d", serviceName, i)
			}
			i++
		}
	}()
	return ch

}

// 优雅退出
func main() {
	doneCh := make(chan struct{})
	c1 := makeGen("service1", doneCh)
	for i := 0; i < 5; i++ {
		select {
		case n := <-c1:
			fmt.Println(n)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
		}
	}
	// 最多不会超过5秒
	doneCh <- struct{}{}
	<-doneCh

}
