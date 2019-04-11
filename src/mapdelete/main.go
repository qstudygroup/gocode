package main

import "fmt"

func main() {
	m := map[string]int{
		"James":            32,
		"Miss Money penny": 27,
	}
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian Flemming")
	fmt.Println(m)

	fmt.Println(m["Miss MoneyPenny"])
	fmt.Println(m["Ian Fleming"])

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m)
}
