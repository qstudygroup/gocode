//lecture 126
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922282?start=0

package main

import (
	"fmt"
)

type circle struct {
	radius float64
}
type square struct {
	length float64
}
type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return c.radius * 3.14 * c.radius
}
func (s square) area() float64 {
	return s.length * s.length
}
func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	c1 := circle{5}
	fmt.Printf("%T\n", &c1)
	info(&c1)

	s1 := square{4}
	info(s1)
}
