package main

import (
	"fmt"
	"math/rand"
	// "sync"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func testStatefulGoroutine() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)

		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("Read:", readOpsFinal)
	fmt.Println("write:", writeOpsFinal)
}

// ============================================
// It seems that the payload of channel are copied.

type Foo struct {
	payload map[int]int
	integer int
}

func testChannelAddress() {
	var send = make(chan Foo)
	var resp = make(chan bool)

	go func() {
		foo := Foo{payload: make(map[int]int), integer: 1}
		fmt.Println("Sender")
		fmt.Printf("Address: %p\n", &foo)
		fmt.Printf("Payload: %v, %p %v\n", foo.payload, foo.payload, foo.integer)
		send <- foo
		foo.payload[1] = 2
		<-resp
		fmt.Println("Sender")
		fmt.Printf("Address: %p\n", &foo)
		fmt.Printf("Payload: %v, %p %v\n", foo.payload, foo.payload, foo.integer)
	}()

	go func() {
		time.Sleep(time.Second)
		foo := <-send
		fmt.Println("Receiver")
		fmt.Printf("Address: %p\n", &foo)
		fmt.Printf("Payload: %v, %p %v\n", foo.payload, foo.payload, foo.integer)
		foo.payload[2] = 2
		foo.integer = 2
		resp <- true

	}()
	time.Sleep(2 * time.Second)
}

func main() {
	// testStatefulGoroutine()
	testChannelAddress()
}
