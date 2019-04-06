package main

import "fmt"

func main() {
	n := "Bond"
	switch n {
	case "Monneypenny", "Bond", "Do No":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("This is default")
	}
}
