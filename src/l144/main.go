// lecture 144
// https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922322?start=0

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(3)
	}()
	fmt.Println(<-ch)
}

func doSomething(x int) int {
	return x * 5
}
