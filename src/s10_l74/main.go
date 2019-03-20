package main

import (
	"fmt"
)

func main() {
	x := []int{12, 34, 435, 54, 32}
	fmt.Println(x[:])
	fmt.Println(x[1:3])

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

}
