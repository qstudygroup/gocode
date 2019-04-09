package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	countLess50 := 0
	countMore50 := 0

	for i := 0; i < 100; i++ {
		x := r.Int63n(100)
		if x < 50 {
			countLess50++
		} else {
			countMore50++
		}
	}
	fmt.Println("countLess50 =", countLess50)
	fmt.Println("CountMore50 =", countMore50)

}
