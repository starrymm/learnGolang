package main

import "fmt"

const(
	BEIJING = iota * 10
	SHANGHAI
	SHENZHEN
)

func main() {

	const lenth int = 10
	fmt.Println("lenght = ", lenth)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)
	
}