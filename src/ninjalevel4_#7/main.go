package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Hellooooooo, James"}
	movie := [][]string{jb, mp}
	fmt.Println(movie)
	for x := range movie {
		fmt.Println(x, movie[x])
		for j, val := range movie {
			fmt.Printf("index position: %v \t value:%v \n", j, val)
		}
	}
}
