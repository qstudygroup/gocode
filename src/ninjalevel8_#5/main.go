// lecture 141
// https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922312?start=0

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (ba byAge) Len() int           { return len(ba) }
func (ba byAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba byAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

type byName []user

func (bn byName) Len() int           { return len(bn) }
func (bn byName) Swap(a, c int)      { bn[a], bn[c] = bn[c], bn[a] }
func (bn byName) Less(a, c int) bool { return bn[a].Last < bn[c].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	sort.Sort(byAge(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, s := range v.Sayings {
			fmt.Println("\t", s)
		}
	}

	fmt.Println("----------------")

	sort.Sort(byName(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, s := range v.Sayings {
			fmt.Println("\t", s)
		}
	}

}
