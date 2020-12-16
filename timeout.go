package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "hello channel"
	}()

	select {
	case v1 := <-ch:
		fmt.Println(v1)
	case <-time.After(2 * time.Second):
		fmt.Println("timed out")
	}
}
