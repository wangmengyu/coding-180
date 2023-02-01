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
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
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
		// 使用非阻塞等待。
		fmt.Println(<-c1)
		select {
		case n := <-c2:
			fmt.Println(n)
		default:
			fmt.Println("not get data from c2")
		}
	}
}
