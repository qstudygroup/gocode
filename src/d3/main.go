//calculate the average of the random numbers from 0 to 100
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0
	numOdds := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		x := r.Intn(100)
		if x%2 == 1 {
			sum += x
			numOdds++
		}
	}
	mean := sum / numOdds
	fmt.Println(mean)
}
