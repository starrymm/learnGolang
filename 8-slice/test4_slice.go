package main

import "fmt"

func main(){
	var numbers = make([]int , 3, 5)

	//向切片追加一个新元素 
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v \n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 3)
	numbers = append(numbers, 4)
	numbers = append(numbers, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v \n", len(numbers), cap(numbers), numbers)

	
}