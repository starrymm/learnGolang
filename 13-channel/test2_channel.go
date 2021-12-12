package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // 有缓冲的channel

	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("goroutine 结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go正在运行, 发送的元素为 = ", i, " len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}

	fmt.Println("main 结束")
}
