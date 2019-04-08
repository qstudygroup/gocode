package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i == 1 {
			fmt.Println(i)
		} else if i == 2 {
			fmt.Println(i)
		} else {
			fmt.Println("it is not 1 or 2")
		}
	}
}
