package main 



import(
	_"GolangStudy/5-init/lib1"  //不调用函数方法， 只会执行init()函数
	mylib2 "GolangStudy/5-init/lib2"  // 起别名，使用别名调函数
	.  "GolangStudy/5-init/lib2"   // 将当前fmt包中的全部方法，导入当前包使用。 不建议使用
)

func main(){
	//lib1.Lib1Test()
	mylib2.Lib2Test()
	//Lib2Test()
}