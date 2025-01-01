package main

import (
	"fmt"
	"runtime"
	"weak"
)

type User struct {
	Age int
}

func main() {
	obj := User{
		Age: 30,
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
