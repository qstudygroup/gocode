//lecture 113
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922254?start=0

package main

import (
	"fmt"
)

func foo(si ...int) int {
	sum := 0
	for _, v := range si {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	fmt.Println(n)

	ee := []int{1, 3, 4, 5, 6, 7, 8, 2, 4}
	m := bar(ee)
	fmt.Println(m)
}
