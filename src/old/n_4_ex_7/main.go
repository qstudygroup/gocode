package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "MoneyPenny", "Helloooooo, James"}
	xy := [][]string{x, y}
	for i, v := range x {

		fmt.Println(i, v)
	}
	for i, v := range y {
		fmt.Println(i, v)
	}
	for i, v := range xy {
		fmt.Println(i, v)
		for j, val := range xy {
			fmt.Printf("index position: %v \t value: \n %v", j, val)
		}

	}
}
