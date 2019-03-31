//lecture 103
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922230?start=0

package main

import "fmt"

func main() {
	x := sum("james", 1, 2, 3, 4)
	fmt.Println("the total is", x)
}

func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("at position", i, "we add", v, "to the total which is now", sum)
	}
	fmt.Println("the total is", sum)
	return sum
}
