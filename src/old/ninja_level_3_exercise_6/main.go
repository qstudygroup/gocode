package main

import "fmt"

func main() {
	a := 2
	if a == 3 {
		fmt.Println("you got your lucky number")
	} else if a < 3 {
		fmt.Println("you didn't get your lucky number and your stupid")
	} else {
		fmt.Println("you are not awful or good")
	}

}
