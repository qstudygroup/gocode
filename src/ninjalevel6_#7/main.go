// lecture 119
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922266?start=0

package main

import (
	"fmt"
)

var g func() = func() {
	fmt.Println("g from outside main said hello")
}

func main() {

	f := func() {
		for i := 0; i < 10; i += 2 {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("this is f %T\n", f)
	g()
	g = f
	g()
	fmt.Printf("this is g %T\n", g)
}
