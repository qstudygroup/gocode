package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func switchOntype(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")

	case string:
		fmt.Println("string")

	case contact:
		fmt.Println("contact")

	default:
		fmt.Println("unknown")
	}
}

func main() {
	switchOntype(7)
	switchOntype("McLeod")
	var t = contact{"Nice to see you", "Tim"}
	switchOntype(t)
	switchOntype(t.greeting)
	switchOntype(t.name)
}
