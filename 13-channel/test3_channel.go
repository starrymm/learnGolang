package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//close关闭一个channel
		close(c)
	}()

	// for {
	// 	//ok为true代表channel没有关闭
	// 	if data, ok := <-c; ok {
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }

	//
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Finished ...")
}
