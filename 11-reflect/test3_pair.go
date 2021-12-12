package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriterBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriterBook() {
	fmt.Println("Writer a Book")
}

func main() {
	b := &Book{}

	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriterBook()

}
