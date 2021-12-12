package main

import (
	"fmt"
)

// interface{}  是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called ....")
	fmt.Println(arg)

	//interface提供类型断言的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T \n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Go"}
	myFunc(book)
	myFunc("hello")
	myFunc(100)
	myFunc(3.14)
}
