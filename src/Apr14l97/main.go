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

	twoPersons := []person{p1, p2}

	for _, p := range twoPersons {
		fmt.Printf("%s %s:", p)
		for _, v := range p1.favIcecreams {
			fmt.Println(v)
		}
	}

}
