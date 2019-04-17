// lecture 142
// https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922316?start=0

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS \t\t", runtime.GOOS)
	fmt.Println("ARCH \t\t", runtime.GOARCH)
	fmt.Println("NumCPU \t\t", runtime.NumCPU())
	fmt.Println("NumGORoutine \t\t", runtime.NumGoroutine())

	wg.Add(1)

	go foo()
	bar()

	fmt.Println("NumCPU \t\t", runtime.NumCPU())
	fmt.Println("NumGORoutine \t\t", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
