package main 

import "fmt"

type myint int 

type Book struct{
	title string
	auth string
}

func checkBook(book Book){
	fmt.Println(book)
	book.title = "c++"
}

func checkBook_(book *Book){
	fmt.Println(*book)
	book.title = "c++"
}

func main()  {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "ws"
	fmt.Println(book1)

	checkBook(book1)
	fmt.Println(book1)

	checkBook_(&book1)
	fmt.Println(book1)
}