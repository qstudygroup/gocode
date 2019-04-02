// lecture 122
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922272?start=0

package main

import (
	"fmt"
)

func main() {
	v := 5
	fmt.Println(v)
	f := enclose()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func enclose() func() int {
	v := 8
	return func() int {
		v++
		return v
	}
}
