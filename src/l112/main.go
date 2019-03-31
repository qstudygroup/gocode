// lecture 112
// https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922250?start=0

package main

import (
	"fmt"
)

func main() {
	n := factorial(5)
	fmt.Println(n)
	m := loopFact(6)
	fmt.Println(m)

}

//usiing recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//using a loop
func loopFact(i int) int {
	total := 1
	for ; i > 0; i-- {
		total *= i
	}
	return total
}
