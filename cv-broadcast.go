package main

import (
	"fmt"
	"sync"
)

var shared = make(map[string]interface{})

func main() {
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	cv := sync.NewCond(&m)

	go func() {
		wg.Add(1)
		defer wg.Done()
		cv.L.Lock()
		for len(shared) == 0 {
			cv.Wait()
		}
		fmt.Printf("GI -> %v\n", shared["one"])
		cv.L.Unlock()
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		cv.L.Lock()
		for len(shared) < 2 {
			cv.Wait()
		}
		fmt.Printf("G2 -> %v\n", shared["two"])
	}()

	cv.L.Lock()
	shared["one"] = "first"
	shared["two"] = "second"
	cv.Broadcast()
	cv.L.Unlock()

	wg.Wait()
}
