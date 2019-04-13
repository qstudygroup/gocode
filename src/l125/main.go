package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x b4", &x)
	fmt.Println("x b4", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y b4", y)
	fmt.Println("y b4", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
