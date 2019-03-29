package main

import (
	"fmt"
)

func main() {
	x := []int{12, 43, 32, 54, 56, 87, 65, 76, 756, 23}
	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", x)
}
