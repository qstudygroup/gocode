package main

import "fmt"

func main() {
	g := (42 == 42)
	h := (42 <= 43)
	i := (42 >= 43)
	j := (42 != 43)
	k := (42 < 43)
	l := (42 > 43)

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}
