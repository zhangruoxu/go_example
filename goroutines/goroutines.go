package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v: %v\n", from, i)
	}
}

func main() {
	f("Direct")
	go f("Goroutine")

	go func(msg string) {
		for i := 0; i < 10; i++ {
			fmt.Println(msg)
		}
	}("going")

	time.Sleep(time.Second)
	fmt.Println("Done.")
}
