package main

import (
	"fmt"
)

func main() {
	m := map[int]string{
		1: "name1",
		2: "name2",
	}

	for i := 3; i <= 10; i++ {
		s := fmt.Sprintf("name%d", i)
		m[i] = s
	}

	fmt.Println(m)

}
