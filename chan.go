package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func(a int, b int) {
		v := a + b
		ch <- v
	}(1, 2)

	r := <-ch
	fmt.Printf("Result: %v\n", r)
}
