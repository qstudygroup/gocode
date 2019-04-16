package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		towtrucks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		towtrucks: []string{
			"monster truck",
			"suvvan",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.towtrucks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}
	for i, val := range s.towtrucks {
		fmt.Println(i, val)
	}
}
