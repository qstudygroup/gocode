package main

import "fmt"

func main() {
	s := 0
	for i := 1; i <= 20; i++ {
		if i%4 == 0 {
			s += i
		}
	}
	fmt.Println(s)
}
