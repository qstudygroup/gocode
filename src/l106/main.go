package main

import (
	"fmt"
)

type hotdog int

type human interface {
	speak()
}
type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- the secret agent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed in to bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed in to bar", h.(secretAgent).first)
	}
	fmt.Println("I was passed in to bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Perry",
			last:  "the Platypus",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	sa1.speak()
	sa2.speak()
	p1.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)

	//conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
