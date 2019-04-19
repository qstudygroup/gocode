// lecture 148
//https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11924334?start=0

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("begin Goroutine:", runtime.NumGoroutine())
	fmt.Println("begin CPU:", runtime.NumCPU())
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("mid Goroutine:", runtime.NumGoroutine())
	fmt.Println("mid CPU:", runtime.NumCPU())
	wg.Wait()
	fmt.Println("about to exit")
	fmt.Println("end Goroutine:", runtime.NumGoroutine())
	fmt.Println("end CPU:", runtime.NumCPU())
}

func foo() {
	fmt.Println("hello from foo")
	wg.Done()
}
func bar() {
	fmt.Println("hello from bar")
	wg.Done()
}
