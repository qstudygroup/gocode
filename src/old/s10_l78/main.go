package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "chocolate", "vanilla", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
