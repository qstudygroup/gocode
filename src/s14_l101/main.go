// lecture 101
// https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922226?start=630


package main

import (
	"fmt"
)
//func (r receiver) identifier(parameteres) (return(r)) {...}
func main() {
	foo()
	bar("James")
	st := woo("Moneypenny")
	fmt.Println(st)
	x,y := mouse("Ian", "Fleming")
	fmt.Print(x)
	fmt.Println(y)
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

func mouse(fn string, ln string) (string, bool){
	a := fmt.Sprintln(fn, ln,`, says "Hello"`)
	b := true
	return a,b
}
 