package main 

import "fmt"

func printMap(cityMap map[string]string) {
	//遍历
	for key,value := range cityMap {
		fmt.Println("key = ", key, "value = ", value)
	}
	cityMap["UK"] = "LONDON"
}

func main(){
	cityMap := make(map[string]string)


	cityMap["China"] = "BEIJING"
	cityMap["USA"] = "NewYork"

	//遍历
	for key,value := range cityMap {
		fmt.Println("key = ", key, "value = ", value)
	}
	
	//删除
	delete(cityMap, "China")

	//修改
	cityMap["USA"] = "DC"

	fmt.Println("--------------")

	//遍历
	for key,value := range cityMap {
		fmt.Println("key = ", key, "value = ", value)
	}
	

	printMap(cityMap)


	fmt.Println("--------------")
	for key,value := range cityMap {
		fmt.Println("key = ", key, "value = ", value)
	}
	
}