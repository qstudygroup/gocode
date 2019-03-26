package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 25
	x[1] = 26
	x[2] = 27
	x[3] = 28
	x[4] = 29
	for i, v := range x {
		fmt.Println(i, v)
		fmt.Printf("%T\n", x)
	}
}
