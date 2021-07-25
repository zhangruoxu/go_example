package main

import (
	"fmt"
	"time"
)

const NumRequest = 5

func testRateLimiting() {
	requests := make(chan int, NumRequest)
	limiter := time.Tick(200 * time.Millisecond)

	for i := 0; i < NumRequest; i++ {
		fmt.Println("Send request", i)
		requests <- i
	}

	close(requests)

	for r := range requests {
		limit := <-limiter
		fmt.Printf("Request %v limit %v, time %v\n", r, limit, time.Now())
	}
}

// ============================================

func testBurstyTimer() {
	testBurstyTimer := make(chan time.Time, 3)
	burstyRequests := make(chan int, NumRequest)

	for i := 0; i < 3; i++ {
		testBurstyTimer <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			testBurstyTimer <- t
		}
	}()

	for i := 0; i < NumRequest; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)

	for r := range burstyRequests {
		limit := <-testBurstyTimer
		fmt.Printf("Request %v limit %v, time %v\n", r, limit, time.Now())
	}
}

func main() {
	// testRateLimiting()
	testBurstyTimer()
}
