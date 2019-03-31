package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for x := 0; x < 100; x++ {
		num := r.Intn(100)
		if num%2 == 1 {
			count++
		}
	}
	fmt.Println("there are", count, "odd number")
}
