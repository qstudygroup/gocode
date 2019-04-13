//lecture 116
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922260?start=0

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v %v and I am %v years old\n", p.first, p.last, p.age)
}

func main() {
	Gina := person{
		first: "Gina",
		last:  "Linetti",
		age:   32,
	}
	Gina.speak()
	Jake := person{
		first: "Jake",
		last:  "Peralta",
		age:   32,
	}
	Jake.speak()
	Holt := person{
		first: "Captain",
		last:  "Holt",
		age:   46,
	}
	Holt.speak()
	Rosa := person{
		first: "Rosa",
		last:  "Diaz",
		age:   26,
	}
	Rosa.speak()
}
