package main

import "fmt"

func main() {
	x := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 26; i++ {
		fmt.Printf("%#U\n", x[i])
		fmt.Printf("%#U\n", x[i])
		fmt.Printf("%#U\n", x[i])
	}
}
