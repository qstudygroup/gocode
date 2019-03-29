package main

import "fmt"

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	x[0] = 42
	x[9] = 999
	x = append(x, 12)
	x = append(x, x...)
	fmt.Println(x)
}
