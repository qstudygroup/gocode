package main

import "fmt"

type person struct {
	firstname              string
	lastname               string
	favoriteicecreamflavor []string
}

func main() {
	p := person{
		firstname:              "Dantdm",
		lastname:               "Minecart",
		favoriteicecreamflavor: []string{"chocolate", "martini", "blueberry"},
	}
	p2 := person{
		firstname:              "Quang",
		lastname:               "Doan",
		favoriteicecreamflavor: []string{"strawberry", "vanilla", "mint chocolate chip"},
	}
	fmt.Println(p.firstname)
	fmt.Println(p.lastname)
	for i, v := range p.favoriteicecreamflavor {
		fmt.Println(i, v)
	}
	fmt.Println(p2.firstname)
	fmt.Println(p2.lastname)
	for i, v := range p2.favoriteicecreamflavor {
		fmt.Println(i, v)
	}
}
