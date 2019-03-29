package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 5, 7, 8, 12, 332, 43, 12, 53, 44, 49, 87}
	for _, v := range a {
		if v%2 == 1 {
			fmt.Println("this is your odd number:", v)
			continue
		} else {
			v++
		}
	}
}
