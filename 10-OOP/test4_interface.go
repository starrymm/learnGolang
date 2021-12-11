package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string 
}

type Cat struct{
	color string
}

func (this * Cat) Sleep()  {
	fmt.Println("Cat is Sleep")
}

func (this * Cat) GetColor()  string {
	return this.color
}


func (this * Cat) GetType()  string {
	return "猫"
}
 

type Dog struct{
	color string
}

func (this * Dog) Sleep()  {
	fmt.Println("Dog is Sleep")
}

func (this * Dog) GetColor()  string {
	return this.color
}


func (this * Dog) GetType()  string {
	return "狗"
}
 


func main()  {
	var animal AnimalIF
	animal = &Cat{"Green"}
	animal.Sleep()

	animal = &Dog{"Yellow"}
	animal.Sleep()

	color :=animal.GetColor()
	fmt.Println("color = ", color)
}


 
