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

	m["Todd"] = 33

	v, ok := m["Barnabas"]
	fmt.Println(v, ok)

	if v, ok := m["Moneypenny"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

	for v, k := range m {
		fmt.Println(v, k)
	}

	xi := []int{4, 5, 6, 7, 8}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
