package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for x := 0; x < 100; x++ {
		num := r.Intn(100)
		switch {
		case num <= 20:
			a++
		case num > 20 && num <= 40:
			b++
		case 40 < num && num <= 60:
			c++
		case 60 < num && num <= 80:
			d++
		case 80 < num && num <= 100:
			e++
		}
	}
	fmt.Printf("there are %v number form 0 to 20 \n", a)
	fmt.Printf("there are %v number form 20 to 40 \n", b)
	fmt.Printf("there are %v number form 40 to 60 \n", c)
	fmt.Printf("there are %v number form 60 to 80 \n", d)
	fmt.Printf("there are %v number form 80 to 90 \n", e)
}
