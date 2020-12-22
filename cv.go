package main

import (
	"fmt"
	"sync"
	_ "time"
)

var shared = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	cv := sync.NewCond(&m)

	wg.Add(1)
	go func() {
		defer wg.Done()
		cv.L.Lock()
		for len(shared) == 0 {
			cv.Wait()
		}
		fmt.Println(shared["hello"])
		cv.L.Unlock()
	}()

	cv.L.Lock()
	shared["hello"] = "world"
	cv.Signal()
	cv.L.Unlock()

	wg.Wait()

}
