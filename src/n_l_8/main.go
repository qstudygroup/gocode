package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case true:
		fmt.Println("this should print")
	}
}
