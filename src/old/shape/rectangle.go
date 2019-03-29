package main

import "fmt"

type rectangle struct {
	width  float64
	height float64
}

func (rec rectangle) area() float64 {
	return rec.width * rec.height
}

func main() {
	rec := rectangle{
		width:  3,
		height: 4,
	}
	fmt.Println("the area of the rectangle is", rec.area())

}
