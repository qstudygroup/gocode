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
	m := map[string]person{
		p1.last:p1,
		p2.last:p2,
	}
	fmt.Println(m)
	
	for k,v := range m{
		fmt.Println(k)
		fmt.Println(v)
	}

	
}	