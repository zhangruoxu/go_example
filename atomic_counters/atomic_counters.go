package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64

	var waitGroup sync.WaitGroup

	for i := 0; i < 50; i++ {
		waitGroup.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
				// ops++
			}
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()

	fmt.Println("Operations", ops)
}
