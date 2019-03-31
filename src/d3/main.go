// sum odd number from 1 to 100
// sum 100 random number from 0 to 100, print out the aveage
// print out how many odd numbers

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0

	for x := 0; x <= 100; x++ {
		s1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		s2 := s1.Intn(100)

		if s2%2 == 1 {
			sum += s2
		}
		fmt.Println(s2)

	}
	mean := sum / 100
	fmt.Println(sum)
	fmt.Println(float64(mean))

}
