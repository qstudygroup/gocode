package main

import "fmt"

var x int

type person struct {
	first string
	last  string
}
type foo int

var y foo

const bar = 42

func main() {
	y = bar
	fmt.Printf("%T", int(y))
	fmt.Printf("%T", bar)
	fmt.Println(bar)
}
