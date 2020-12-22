package main

import (
	"fmt"
	//"golang.org/x/net/context"
	"context"
)

func main() {

	generator := func(ctx context.Context) <-chan int {
		ch := make(chan int)

		n := 1
		go func() {
			defer close(ch)
			for {
				select {
				case ch <- n:
				case <-ctx.Done():
					return
				}
				n++
			}
		}()
		return ch
	}

	ctx, cancel := context.WithCancel(context.Background())
	ch := generator(ctx)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			cancel()
		}
	}
}
