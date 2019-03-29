package main

import "fmt"

func main() {
	m := map[string]int{
		"James":            32,
		"Miss Money Penny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)
}
