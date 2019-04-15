package main

import "fmt"

type person struct {
	firstName    string
	lastName     string
	favIcecreams []string
}

func main() {
	p1 := person{
		firstName:    "Quang",
		lastName:     "Doan",
		favIcecreams: []string{"Rainbow", "Yogurt berries", "tonight show"},
	}
	p2 := person{
		firstName:    "everyone",
		lastName:     "mystery",
		favIcecreams: []string{"the everything flavor"},
	}
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for _, val := range v.favIcecreams {
			fmt.Println(val)
		}
		fmt.Println("-----")
	}
}
