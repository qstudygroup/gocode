//lecture 118
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922264?start=0

package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("this is an anonymous func.")
	}()

	func() {
		for i := 0; i < 100; i += 2 {
			fmt.Println(i)
		}
	}()
}
