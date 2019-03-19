package main

import "fmt"

func main() {
	favsport := "soccer"
	switch favsport {

	case "swimming":
		fmt.Println("go to the pool")
	case "soccer":
		fmt.Println("go to the soccer field")
	case "skiing":
		fmt.Println("go to a mountain")

	}
}
