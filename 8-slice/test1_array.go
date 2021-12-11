package main

import "fmt"

func printArray(myArray [10]int){
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}

	
}

func main() {
	var myArray1 [10]int;
	myArray2 := [10]int{1,2,3,4}
	//myArray3 := [4]int{1,2,3,4}
	for i := 0; i < 10; i++ {
		fmt.Println(myArray1[1])
	}
	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}


	//查看数组数据类型
	fmt.Printf("myArray1 types = %T \n", myArray1)
	fmt.Printf("myArray1 types = %T \n", myArray2)

	printArray(myArray1)

}