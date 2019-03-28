package main

import (
	"fmt"
)

func main() {
	foo()
	bar("James")
	st := woo("Moneypenny")
	fmt.Println(st)
}

func foo() {
	fmt.Println("hello")
}

func bar(s string) {
	fmt.Println("hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("hello from woo, ", st)
}
