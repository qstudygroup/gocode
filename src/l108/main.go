//lecture 108
// https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922240?start=0

package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("my first expression")
	}
	f()

	g := func(x int) {
		fmt.Println("the year big brother started watching:", x)
	}
	g(1984)
}
