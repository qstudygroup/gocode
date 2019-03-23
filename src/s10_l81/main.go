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
	delete(m, "James")
	fmt.Println(m)
	delete(m, "Ian Fleming")
	fmt.Println(m["Moneypenny"])
	fmt.Println(m["Ian Fleming"])

	if v, ok := m["Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Moneypenny")
	}
	fmt.Println(m)
}
