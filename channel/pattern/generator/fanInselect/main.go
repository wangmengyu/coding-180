package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 同样也是等待多个通道内的数据两边一起接受处理掉
func fanInBySelect(c1, c2 chan string) chan string {
	ch := make(chan string)
	go func() { // goroutine 写一个就可以了. 用select进行区分接受
		for {
			select {
			case v := <-c1:
				ch <- v
			case v := <-c2:
				ch <- v
			}
		}

	}()

	return ch

}

// 不断的从管道内生成数字.
func makeGen(serviceName string) chan string {
	c := make(chan string)
	go func() {
		// 在管道内生成. 即goroutine
		v := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("%s receive %d", serviceName, v)
			v++
		}
	}()
	return c
}
func main() {

	c1 := makeGen("service1")
	c2 := makeGen("service2")
	ch := fanInBySelect(c1, c2) //

	// 提取所有总通道里返回的数据
	for v := range ch {
		fmt.Println(v)
	}

}
