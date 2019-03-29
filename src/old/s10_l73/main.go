package main

import (
	"fmt"
)

func main() {
	x := []int{12, 34, 435, 54, 32}
	fmt.Println(x)
	fmt.Println(len(x))

	for i, v := range x {
		fmt.Println(i, v)
	}

}
