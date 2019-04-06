// lecture 131
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922294?start=0

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First`
	Last  string `json:"Last`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Miss","Last":"Moneypenny","Age":27},{"First":"James","Last":"Bond","Age":32}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("error:", err)
	}
	for i, v := range people {
		fmt.Println("person number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
