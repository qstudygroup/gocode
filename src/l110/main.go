// lecture 110
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922244?start=0

package main

import (
	"fmt"
)

func sum(si ...int) int {
	fmt.Printf("%T\n", si)
	total := 0
	for _, v := range si {
		total += v
	}
	return total
}

func even(f func(si ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(si ...int) int, vi ...int) int {
	var ei []int
	for _, v := range vi {
		if v%2 == 1 {
			ei = append(ei, v)
		}
	}
	return f(ei...)
}

func main() {
	ii := []int{1, 2, 3, 4, 6, 7, 5, 75, 3, 3, 86, 8}
	fmt.Println(sum(ii...))

	e := even(sum, ii...)
	fmt.Println("even numbers", e)

	o := odd(sum, ii...)
	fmt.Println("odd numbers", o)
}
