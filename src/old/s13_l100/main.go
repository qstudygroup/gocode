package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		dog           bool
		possibleBreed map[string]string
	}{
		dog: true,
		possibleBreed: map[string]string{
			"Golden Retriver": "female",
			"German Sheperd":  "male",
		},
	}
	fmt.Println(p1)
	for i, v := range p1.possibleBreed {
		fmt.Println(i, v)
	}
}
