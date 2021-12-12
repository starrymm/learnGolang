package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called ..")
	fmt.Printf("%v \n", this)
}

func main() {
	user := User{1, "mm", 17}

	DoFiledAndMethod(user)

}

func DoFiledAndMethod(input interface{}) {

	//通过input获取type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	//input获取alue
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	//
	for i := 0; i < inputType.NumField(); i++ {
		filed := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s, %v = %v\n", filed.Name, filed.Type, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s, %v\n", m.Name, m.Type)
	}
}
