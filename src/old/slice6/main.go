package main

import "fmt"

func main() {
	x := []int{4, 5, 8, 42}
	x = append(x, 7, 328192983, 93939, 83838)
	fmt.Println(x)
	y := []int{452, 789, 987, 1234}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
