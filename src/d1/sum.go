package main

import "fmt"

func main() {
	s := 0
	for x := 1; x <= 100; x = x + 2 {
		s = s + x
	}
	fmt.Println(s)
}
