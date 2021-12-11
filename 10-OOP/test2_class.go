package main 

import "fmt"


//如果类名首字母大写，表示其它包报能够访问
type Hero struct{
	//如果首字母大写，该属性对外可访问，否则的话只有类内部能访问
	Name string
	Ad int 
	level int
}

// func (this Hero) Show() {
// 	fmt.Println("Name = ", this.Name)
// 	fmt.Println("Ad = ", this.Ad)
// }

// func (this Hero) GetName() string {
// 	return this.Name
// }

// func (this Hero) SetName(newName string){
// 	//this 是对象的副本
// 	this.Name = newName
// }

func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string){
	//this 是对象的副本
	this.Name = newName
}

func main() {

	//创建一个对象
	hero := Hero{Name: "zhangs", Ad:100, level: 1}
	//hero.GetName()
	hero.Show()
	hero.SetName("li4")
	hero.Show()

}
