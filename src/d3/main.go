package main

import (
	"fmt"
	"math/rand"
	"time"
)

type friend struct {
	money int
	grade int
}

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	var friends []friend
	for i := 0; i <= 99; i++ {
		g := r.Intn(10)
		m := r.Intn(1000)
		f := friend{
			money: m,
			grade: g,
		}
		friends = append(friends, f)
	}
	for i, f := range friends {
		fmt.Println(i, f.money, f.grade)
	}
}
