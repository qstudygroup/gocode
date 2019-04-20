// lecture 150
// https://www.udemy.com/learn-how-to-code/learn/v4/t/lecture/11922334?start=0

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
			fmt.Println(counter)
		}()
	}
	wg.Wait()
	fmt.Println("end counter", counter)
}
