package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================
// Test worker group.

// Start n workers. Obtain jobs from the jobs channel
// and send results to the results channel.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "starting job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}
}

func workerGroups() {
	const numWorkers = 3
	const numJobs = 5

	jobs := make(chan int)
	results := make(chan int)

	// Start n workers.
	for w := 0; w < numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Start a new goroutine for sending jobs.
	// Since workers are blocked on obtaining jobs from the channel,
	// all worker goroutines are sleeping. Then the runtime complains that:
	//
	// fatal error: all goroutines are asleep - deadlock!
	//
	// But the main thread are still living.
	// This is quite wierd.
	// Therefore, we use a goroutine to send jobs.
	go func() {
		for j:= 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Obtain rsults from the results channel.
	for a := 1; a < numJobs; a++ {
		<-results
	}
}

// ============================================

func waitGroupWorker(id int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	fmt.Printf("Worker %d is staring.\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d ends.\n", id)
}

func waitGroups() {
	const numWorkers = 5
	var waitGroup sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		waitGroup.Add(1)
		go waitGroupWorker(i, &waitGroup)
	}

	waitGroup.Wait()
}

func main() {
	// workerGroups()
	waitGroups()
}