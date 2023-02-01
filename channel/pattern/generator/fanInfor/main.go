package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 任意多的输入管道时适用该方法
func fanIn(chs ...chan string) chan string {

	c := make(chan string)
	for _, ch := range chs {
		// 每个输入管道给予goroutine

		go func(ch chan string) { // 这里一定要用函数传参才能正确拷贝一份ch, 否则可能用了不对的ch
			for {
				c <- <-ch // for 内部的 goroutine 如果调用 for 的 循环变量 可能不是希望的！  所以要额外复制。
			}
		}(ch)
	}

	return c
}

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
	ch1 := makeGen("service1") // 相当一个服务, 拿着它就能和其他服务进行交互了
	ch2 := makeGen("service2") // 相当一个服务, 拿着它就能和其他服务进行交互了
	ch3 := makeGen("service3") // 相当一个服务, 拿着它就能和其他服务进行交互了
	c := fanIn(ch1, ch2, ch3)

	// 读取 c 来源的数据
	for n := range c {
		fmt.Println(n)
	}
}
