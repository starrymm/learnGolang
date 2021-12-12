package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		defer fmt.Println("A defer")

		func() {
			defer fmt.Println("B defer")
			//推出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	go func(a int, b int) bool {
		fmt.Println("a = ", a, "b = ", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
