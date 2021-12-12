package main

import (
	"fmt"
	"time"
)

//从goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}

}

//主goroutine
func main() {

	// 创建一个go程去执行newTask()
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main Goroutine: i = %d \n", i)
		time.Sleep(1 * time.Second)
	}

}
