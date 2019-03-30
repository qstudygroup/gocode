//lecture 107
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922238?start=122

package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("func anonymous ran")
	}()
	func(x int) {
		fmt.Println("x =", x)
	}(42)
}

func foo() {
	fmt.Println("Func foo ran")
}
