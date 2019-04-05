//lecture 129
// https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922290?start=0

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//JSON Marshal
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
