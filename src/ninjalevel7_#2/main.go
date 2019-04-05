//lecture 128
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922286?start=0

package main

import (
	"fmt"
)

type person struct {
	age  int
	name string
}

func changeMe(p *person) {
	(*p).name = "Miss Moneypenny"
}

func main() {
	p1 := person{
		age:  23,
		name: "James",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
