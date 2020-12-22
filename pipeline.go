package main

import (
	"fmt"
)

func generator(nums ...int) <-chan int {
	r := make(chan int)
	go func() {
		for _, n := range nums {
			r <- n
		}
		close(r)
	}()
	return r
}

func square(ch <-chan int) <-chan int {
	s := make(chan int)
	go func() {
		for v := range ch {
			s <- v * v
		}
		close(s)
	}()
	return s
}

func main() {
	for v := range square(generator(2, 3)) {
		fmt.Println(v)
	}
}
