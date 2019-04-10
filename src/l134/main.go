//lecture 134
// https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922300?start=0

package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)

}
