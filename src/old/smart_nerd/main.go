package main

import (
	"fmt"
)

func main() {
	switch "Bond" {

	case "MoneyPenny":
		fmt.Println("Miss Money")
	case "Bond":
		fmt.Println("Bond James")
	case "Q":
		fmt.Println("this is Q")
	default:
		fmt.Println("this is default")
	}
}
