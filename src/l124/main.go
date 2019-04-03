package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // & gives you the address

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value store at an address
	fmt.Printf("%T\n", b)

	*b = 43
	fmt.Println(a)

}
