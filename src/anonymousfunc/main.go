package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()
	func(x int) {
		fmt.Println("the meaning life:", x)
	}(42)
}
func foo() {
	fmt.Println("foo ran")
}
