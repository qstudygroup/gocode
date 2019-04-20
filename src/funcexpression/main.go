package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")

	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("the year big brother started watching")
	}
	g(1984)
}
