// lecture 149
// https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("I am %v%v", p.first, p.last)
}
func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "James ",
		last:  "Bond",
	}
	saySomething(&p)
	//saySomething(p)
}
