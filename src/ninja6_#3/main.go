//lecture 115
// https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922258?start=0
package main

import (
	"fmt"
)

func main() {
	foo()
	defer foo()
	bar()
}

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}

func bar() {
	fmt.Println("bar ran")
}
