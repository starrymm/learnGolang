package main

import "fmt"

func main() {

	var a string
	//pair<statictype:string, value:"abcd">
	a = "abcd"

	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
