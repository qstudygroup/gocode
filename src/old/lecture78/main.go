package main

import(
	"fmt"
)
func main(){
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}       

	fmt.Println(mp)

	xp := [][]string{mp, jb}
	fmt.Println(xp)
}