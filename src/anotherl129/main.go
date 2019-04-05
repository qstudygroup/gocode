//lecture 129
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922290?start=0

package main

import (
	"encoding/json"
	"fmt"
)

var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll", "Order": "Dasyuromorphia"}
	]`)

type Animal struct {
	Name  string
	Order string
}

var animals []Animal

func main() {
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
