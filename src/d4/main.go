package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 0
	for x := 0; x <= 100; x++ {
		s1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		s2 := s1.Intn(100)

		if s2%2 == 1 {
			count++
		}
		fmt.Println(s2)
	}
	fmt.Println("there are", count, "odd number")

}
