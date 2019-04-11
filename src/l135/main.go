//lecture 135
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922302?start=0

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginWord1 := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginWord1))
	if err != nil {
		fmt.Println("YOU CAN'T LOG IN")
		return
	}
	fmt.Println("You're logged in")
}
