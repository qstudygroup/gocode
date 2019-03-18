package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const myLuckyNumber = 10
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numBadLucks := 0

	for {
		n := r.Intn(11)
		if n == myLuckyNumber {
			break
		} else {
			numBadLucks++

		}
	}
	fmt.Println("number of bad lucks =", numBadLucks)
}
