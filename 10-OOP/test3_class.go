package main

import "fmt"

type Human struct {
	name string
	sex string
}

func (this * Human) Eat()  {
	fmt.Println("Human .Eat()...")
}

func (this * Human) Walk()  {
	fmt.Println("Human .Walk()...")
}

type SuperMan struct{
	Human // SuperMan类继承了Human方法
	level int 
}


//重写父类的方法Eat()
func (this *SuperMan) Eat(){
	fmt.Println("SuperMan.Eat() ...")
}


//字类的新方法
func (this *SuperMan) Fly(){
	fmt.Println("SuperMan.Fly() ...")
}

func (this *SuperMan) Show()  {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main()  {
	h := Human{"zs", "felman"}

	h.Eat()
	h.Walk()

	//定义子类对象
	s := SuperMan{Human{"li4", "female"}, 80}

	s.Walk()
	s.Eat()
	s.Fly()

	var s1 SuperMan
	s1.name = "li2"
	s1.sex = "male"
	s1.level = 65
	s1.Show()
}