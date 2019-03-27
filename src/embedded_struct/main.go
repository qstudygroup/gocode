package main

import "fmt"

var x int
var y int

func main() {
	x := 42
	y := 43
	fmt.Println(x, y)

	x, y = 44, 45
	fmt.Println(x, y)
}
