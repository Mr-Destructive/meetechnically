package main

import (
	"fmt"
	"runtime"
	"weak"
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
	ref2 := ref
	ref3 := weak.Make(ref)
	fmt.Println("Original Object:", obj)
	fmt.Println("Original Reference:", ref)
	fmt.Println("Second Reference:", ref2)
	fmt.Println("Third Reference:", ref3)

	ref = nil
	ref2 = nil
	runtime.GC()
	fmt.Println("Original Reference:", ref)
	fmt.Println("Second Reference:", ref2)
	fmt.Println("Third Reference:", ref3.Value())
}
