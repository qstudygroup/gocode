// lecture 113
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922252?start=0

package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	n := 4
	return n
}

func bar() (int, string) {
	n := 5
	m := "hello"
	return n, m
}
