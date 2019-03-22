package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])

	v, ok := m["Barnabas"]
	fmt.Println(v, ok)

	if v, ok := m["Moneypenny"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}
}
