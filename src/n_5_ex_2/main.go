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
	m := map[string]person{
		p.lastname:  p,
		p2.lastname: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.firstname)
		fmt.Println(v.lastname)
		for i, val := range v.favoriteicecreamflavor {
			fmt.Println(i, val)

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
	}
}
