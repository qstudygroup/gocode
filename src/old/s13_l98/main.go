package main

import (
	"fmt"
)

type person struct {
	first           string
	last            string
	iceCreamFlavour []string
}

func main() {
	p1 := person{
		first:           "Quang",
		last:            "Doan",
		iceCreamFlavour: []string{"Lemon Sorbet", "Vanilla"},
	}
	p2 := person{
		first:           "Giang",
		last:            "Nguyen",
		iceCreamFlavour: []string{"Chocolate", "Cookie n Cream"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.iceCreamFlavour {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}

}
