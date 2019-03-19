package main

import (
	"fmt"
)

func main() {
	favSport := "swim"
	switch favSport {
	case "fly":
		fmt.Println("i cant fly")
	case "swim":
		fmt.Println("go to the pool")
	case "jog":
		fmt.Println("go out")
	}
}
