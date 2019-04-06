// lecture 130
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922292?start=0

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}
	p2 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	people := []person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bs))

}
