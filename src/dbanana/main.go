package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 5, 7, 8}
	for i := 0; i < len(a); i++ {
		v := a[i]
		if v%2 == 0 {

			fmt.Println(v)
		}
	}
}
