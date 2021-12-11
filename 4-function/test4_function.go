package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c 
}

// 返回多个返回之，匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 0, 1
}


//返回多个返回值， 有星灿名称
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---foo2---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 1000
	r2 = 1000
	return 
}



func main(){
	c  := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("aa", 999)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)


	ret1, ret3 := foo3("hh", 22)
	fmt.Println("ret1 = ", ret1, "ret3 = ", ret3)

}