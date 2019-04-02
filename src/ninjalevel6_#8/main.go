//lecture 120
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922268?start=0

package main

import (
	"fmt"
)

func foo() func() int {
	return func() int {
		return 42
	}
}

func main() {
	a := foo()
	fmt.Println(a())

}
