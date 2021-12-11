package main 

import "fmt"

func main()  {
	
	//第一种形式
	//key:string value:string
	var myMap1 map[string]string

	if(myMap1 == nil) {
		fmt.Println("myMap 为空")
	}
	//给map开辟空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "C++"
	myMap1["three"] = "Python"
	fmt.Println(myMap1)


	//第二种形式
	myMap2 := make(map[string]string)
	myMap2["one"] = "java"
	myMap2["two"] = "C++"
	myMap2["three"] = "Python"
	fmt.Println(myMap2)

	//第三种形式
	myMap3 := map[string]string{
		"one": "php",
		"two": "c++",
	}
	fmt.Println(myMap3)

}