package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {

		//如果c可写
		case c <- x:
			tmp := x
			x = y
			y = y + tmp

		//如果c可读
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibonacii(c, quit)
}
