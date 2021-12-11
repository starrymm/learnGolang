package main

import(
	"fmt"
)

func main(){
	var a int;
	fmt.Println("a = ", a);

	var b int = 100;
	fmt.Println("b = ", b);

	var c = 100;
	fmt.Println("c = ", c);
	fmt.Printf(" type of c = %T\n", c);

	var bb = "abcd";
	fmt.Printf("bb = %s type of bb = $T \n", bb, bb);

	e := 100
	fmt.Println("e = ", e);
	fmt.Printf("type of f = %T \n", e);

	var xx, yy int = 100, 200;
	fmt.Println("xx = ", xx, ", yy = ", yy)

	var (
		vv int = 100;
		jj bool = true;
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)

	

}