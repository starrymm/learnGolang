package main

import "fmt"

func main(){

	//写入defer 关键字
	defer fmt.Println("main end 1")
	defer fmt.Println("main end 2")

	fmt.Println("main hello world 1")
	fmt.Println("main hello world 2")
}