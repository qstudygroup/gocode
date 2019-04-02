// lecture 122
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922272?start=0

package main

import (
	"fmt"
)

func main() {
	v := 5
	fmt.Println(v)
	enclose()
}

func enclose() {
	v := 8
	fmt.Println(v)
	v++
	fmt.Println(v)
}
