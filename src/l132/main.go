//lecture 132
// writer  interface

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello")
	fmt.Fprint(os.Stdout, "hello")
	io.WriteString(os.Stdout, "hello")
}
