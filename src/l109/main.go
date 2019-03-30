// lecture 109
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922242?start=0

package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	fi := bar()
	fmt.Printf("%T\n", fi)

	fmt.Println(fi())
	fmt.Println(bar()())

}

// return a string
func foo() string {
	return "Hello world"
}

// return a func
func bar() func() int {
	return func() int {
		return 451
	}
}
