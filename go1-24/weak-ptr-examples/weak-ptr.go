package main

import (
	"fmt"
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
	weakRef := weak.Make(ref)
	fmt.Println("Original Object:", obj)
	fmt.Println("Strong Reference:", ref)
	fmt.Println("Weak Reference:", weakRef)

	if weakRef.Value() != nil {
		fmt.Println("Object is not yet garbage collected!", weakRef.Value())
	} else {
		fmt.Println("Object is cleaned by the garbage collector")
	}

	if strongRef := weakRef.Value(); strongRef != nil {
		fmt.Println("Object is not yet garbage collected!", strongRef)
	} else {
		fmt.Println("Object is cleaned by the garbage collector")
	}
}
