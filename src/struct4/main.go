package main

import "fmt"

type person struct {
	first                  string
	last                   string
	favoriteIceCreamFlavor []string
}

func main() {
	p1 := person{
		first:                  "Nhim",
		last:                   "The great",
		favoriteIceCreamFlavor: []string{"Chocolate"},
	}
	p2 := person{
		first:                  "Quang",
		last:                   "Doan",
		favoriteIceCreamFlavor: []string{"lemon sorbet"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favoriteIceCreamFlavor {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favoriteIceCreamFlavor {
		fmt.Println(i, v)
	}
}
