package main

import (
	"fmt"
)

func main() {
	x := sum("James")
	fmt.Println("the total is", x)
}

func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, "we are now adding,", v, "to the total which is now,", sum)
	}
	fmt.Println("the total is,", sum)
	return sum
}
