package main

import "fmt"

func main() {
	jb := []string{"James ", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Money", "Penny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	xp := [][]string{mp, jb}
	fmt.Println(xp)
}
