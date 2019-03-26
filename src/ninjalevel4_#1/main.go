package main

import (
	"fmt"
)

func main() {
	var x [5]int
	x[0] = 2
	x[1] = 32
	x[2] = 90
	x[3] = 98
	x[4] = 76
	for i, v := range x {
		fmt.Println(i, v)
		fmt.Printf("%T\n", x[i])
	}
}
