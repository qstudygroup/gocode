package main

import "fmt"

func main() {
	m := map[string]int{
		"James":            32,
		"Miss money Penny": 27,
	}
	fmt.Println(m)
	delete(m, "James")
	fmt.Println(m)
	delete(m, "Ian Flemming")
	fmt.Println(m)

	fmt.Println(m)

	fmt.Println(m["Miss money Penny"])

	fmt.Println(m["Ian Flemming"])
	if v, ok := m["Miss money Penny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss money Penny")

	}
	fmt.Println(m)

}
