package main

import (
	"fmt"
	"time"
)

func testChannels() {
	messages := make(chan string)

	go func() {
		messages <- "A piece of message."
		messages <- "Another piece of message."
	}()

	msg := <-messages
	fmt.Println("Message:", msg)
	msg = <-messages
	fmt.Println("Message:", msg)
}

// ============================================

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(5 * time.Second)
	fmt.Println("Done")

	done <- true
}

func testChannelSync() {
	done := make(chan bool)
	go worker(done)
	fmt.Println("Wating...")
	<-done
	fmt.Println("main finished.")
}

// ============================================

func ping(pings chan<- string, message string) {
	pings <- message
}

func pong(pings <-chan string, pongs chan<- string) {
	message := <-pings
	// The following statement won't compile:
	//
	// pings <- "test"
	//
	// since pings are declared as receiving,
	// sending messages to a channel which declared as receiving
	// will cause compilation error.
	pongs <- message
}

func testChannelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	message := "Ping pong ball"

	ping(pings, message)
	pong(pings, pongs)
	fmt.Println("Message:", <-pongs)
}

// ============================================

func testSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// ============================================
// Channel sending and receiving are synchronous operations.

func testChannelBlocking() {
	c := make(chan string)
	go func() {
		fmt.Println("Before sending...")
		c <- "A piece of message."
		fmt.Println("Sending finished.")
	}()

	fmt.Println("Before receiving...")
	time.Sleep(5 * time.Second)
	fmt.Println("Message:", <-c)
	time.Sleep(1 * time.Second)
}

// ============================================
// Buffered channels can achieve non-blocking sending and receiving.
// It seems that if the buffer of a channel overflows,
// i.e., the number of messages is larger than the size of channel buffer,
// the channel beccomes a synchronous one again.

func testChannelNonBlocking() {
	c := make(chan string, 2)
	go func() {
		fmt.Println("Before sending...")
		c <- "Message one."
		c <- "Message two."
		c <- "Message three."
		fmt.Println("Sending finished.")
	}()

	fmt.Println("Before receiving...")
	time.Sleep(5 * time.Second)
	fmt.Println("Message:", <-c)
	fmt.Println("Message:", <-c)
	time.Sleep(1 * time.Second)
}

// ============================================

func testChannelTimeout() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Result 1"
	}()

	select {
	case result := <-c1:
		fmt.Println("Result:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Result 2"
	}()

	select {
	case result := <-c2:
		fmt.Println("Result:", result)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout 2")
	}
}

// ============================================

func testClosingChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Receiving job", j)
			} else {
				fmt.Println("Received all jobs")
				done <- true
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("Sending job", i)
	}

	close(jobs)
	fmt.Println("All jobs are sent.")

	<-done
}

// ============================================

func testRangeOverChannel() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)
	fmt.Println("Channel closed.")

	for s := range queue {
		fmt.Println(s)
	}
}

// ============================================

func testTimers() {
	timer1 := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting for timer 1 ...")
	val := <-timer1.C
	fmt.Println("Timer 1 fired", val)

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired.")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped.")
	}
}

// ============================================

func testTickers() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped.")
}

// ============================================
// In Java, the termination of the main thread will not cause the termination
// of other threads spawned from the main thread. Unless one user-created
// Java threads are running, the Java process will not exit.

// In Golang, the main thread terminates then all the running goroutines terminate.
// I have no idea why it is.

func testWhenGoroutineStop() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Child", i)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("Main finished.")
}

// ============================================
// Understanding channels and bufferred channels.
// Sending messages to channels seems a synchronous operation:
// sending statement is block until there is a goroutine reads from the channel.
// Consequently, Go provides buffered channels, where the number of buffered messages
// is specified. Sending messages not exceeding the number of buffered messages
// will not block the sending statement.

func testBufferedChannels() {
	msg := make(chan string, 1)
	done := make(chan bool)

	go func() {
		msg <- "Message 1"
		fmt.Println("Message 1 sent.")
		msg <- "Message 2"
		fmt.Println("Message 2 sent.")
		done <- true
	}()

	go func() {
		fmt.Println("Sleeping...")
		time.Sleep(5 * time.Second)
		fmt.Println(<- msg)
		done <- true
	}()

	<- done
	<- done
}

func main() {
	// testChannels()
	// testChannelSync()
	// testChannelDirections()
	// testSelect()
	// testChannelWait()
	// testChannelBlocking()
	// testChannelNonBlocking()
	// testChannelTimeout()
	// testClosingChannel()
	// testRangeOverChannel()
	// testTimers()
	// testWhenGoroutineStop()
	// testTickers()
	testBufferedChannels()
}
