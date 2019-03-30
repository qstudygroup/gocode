package main

import (
	"fmt"
)

func main() {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len(s); i++ {
		for v := 0; v <= 2; v++ {
			fmt.Printf("%#U\n", s[i])
		}
		fmt.Println("")
	}

}
