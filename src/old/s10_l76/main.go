package main

import (
	"fmt"
)

func main() {
	x := []int{12, 34, 435, 54, 32}
	fmt.Println(x)
	x = append(x, 23, 53, 64, 7532)
	fmt.Println(x)

	y := []int{2434, 76543, 132, 75654}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
