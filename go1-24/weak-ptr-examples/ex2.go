package main

import (
	"fmt"
	"runtime"
	"weak"
)

//type User struct {
//   Name string
//	Age int
//	Email string
//}

func main() {
	obj := User{
		Name:  "John",
		Age:   30,
		Email: "john@mail.com",
	}

	ref := &obj
	weakPtr := weak.Make(ref)

	fmt.Println("Object:", obj)
	fmt.Println("Reference:", ref)
	fmt.Println("Weak pointer:", weakPtr)

	ref = nil
	runtime.GC()

	fmt.Println("Reference:", ref)
	fmt.Println("Weak pointer:", weakPtr)
	fmt.Println("Weak pointer reference:", weakPtr.Value())
	if strongPtr := weakPtr.Value(); strongPtr != nil {
		fmt.Println("Object is not yet garbage collected!", strongPtr)
	} else {
		fmt.Println("Object is cleaned by the garbage collector")
	}

}
