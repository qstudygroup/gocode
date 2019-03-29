package main

import (
	"fmt"
)

func main() {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
		fmt.Printf("%#U\n", s[i])
		fmt.Printf("%#U\n\n", s[i])
	}
}
