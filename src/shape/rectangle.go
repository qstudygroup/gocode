package main

import (
	"fmt"
)

type rectangle struct {
	width  float64
	height float64
}

func (rec rectangle) perimeter() float64 {
	return (rec.width + rec.height) * 2
}
func main() {
	rec := rectangle{
		width:  3,
		height: 4,
	}
	fmt.Println("the perimeter of the rectangle is:", rec.perimeter())
}
