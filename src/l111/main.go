//lecture 111
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922246?start=0

package main

import (
	"fmt"
)

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}

}
