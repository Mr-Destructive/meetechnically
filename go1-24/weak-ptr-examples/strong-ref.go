package main

import (
	"fmt"
	"runtime"
)

//type User struct {
//	Name  string
//	Age   int
//	Email string
//}

func main() {
	obj := User{
		Name:  "John",
		Age:   30,
		Email: "john@mail.com",
	}
	ref := &obj
	fmt.Println("Original Object:", obj)
	fmt.Println("Original Reference:", ref)

	ref = nil
	runtime.GC()
	fmt.Println("Original Reference:", ref)
}
