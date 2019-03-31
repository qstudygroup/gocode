package main

import "fmt"

func main() {
	s := "fortnite is awesome"
	fmt.Println(s)
	s = "Hello, Hawaii"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println("")

	for i, v := range s {
		fmt.Printf("at index positon %d we have hex %#x\n", i, v)
	}
}
