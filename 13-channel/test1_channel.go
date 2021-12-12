package main

import "fmt"

func main() {

	//定义一个channel

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine 运行...")

		c <- 666
	}()

	num := <-c //从c中读取数据

	fmt.Println("num = ", num)
	fmt.Println("main goroutine结束")
}
