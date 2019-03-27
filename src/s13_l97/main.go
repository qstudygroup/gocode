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

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.iceCreamFlavour {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first, p2.last)
	for i, v := range p2.iceCreamFlavour {
		fmt.Println(i, v)
	}
}
