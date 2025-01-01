package main

import (
	"fmt"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {

	obj := User{
		Age: 30,
	}

	ref := &obj

	fmt.Println("Object:", obj)
	fmt.Println("Reference:", ref)

	fmt.Printf("Objects's Address: %p\n", &obj)
	fmt.Printf("Reference's Address: %p\n", ref)
}
