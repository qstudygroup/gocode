package main

import(
	"fmt"
)

func main(){
	s := struct{
		first string
		friends map[string]int
		favdrinks[]string

	}{
		first: "James",
		friends: map[string]int{
			"Monneypenny": 555,
			"Q": 777,
			"M": 888,
		},
		favdrinks: []string{
			"Martini",
			"water",
		},
	}
		fmt.Println(s.first)
		fmt.Println(s.friends)
		fmt.Println(s.favdrinks)
		for k, v := range s.friends{
			fmt.Println(k, v)
		}
		for i, val := range s.favdrinks{
			fmt.Println(i, val)
		}
}
