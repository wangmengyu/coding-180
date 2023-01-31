package main

import "fmt"

func chanDemo() {
	//定义一个channel
	ch := make(chan struct{})
	// 往里面推送数据
	ch <- struct{}{}
	ch <- struct{}{}
	// 尝试接收一个数据
	v1 := <-ch
	// 没有任何协程接收管道的数据。
	fmt.Println(v1)
}
func main() {
	chanDemo()

}
