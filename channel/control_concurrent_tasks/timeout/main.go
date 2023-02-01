package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 返回一个能不断生成字符串的管道
func makeGen(serviceName string) chan string {
	ch := make(chan string)
	// 生成数据需要在独立的协程里完成

	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
			ch <- fmt.Sprintf("%s receive %d", serviceName, i)
			i++
		}
	}()
	return ch

}

func main() {
	c1 := makeGen("service1")
	c2 := makeGen("service2")

	for {
		fmt.Println(<-c2)
		select {
		case v := <-c1:
			fmt.Println(v)
		case <-time.After(2000 * time.Millisecond):
			fmt.Println("timeout ")
		}

	}

}
