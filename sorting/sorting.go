package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// ============================================
// Passing a slice to the sort.Strings().

func sorting() {
	strs := [...]string{"a", "c", "b"}
	fmt.Println("Before soring:", strs)
	sort.Strings(strs[:])
	fmt.Println("After soring:", strs)
}

// ============================================
// Sort large integer arrays.

const LENGTH = 10000000

func genLargeArray() []int {
	ints := new([LENGTH]int)
	for i := 0; i < len(ints); i++ {
		ints[i] = rand.Intn(LENGTH)
	}
	return ints[:]
}

func sortLargeArray() {
	ints := genLargeArray()
	fmt.Println("Is sorted:", sort.IntsAreSorted(ints))
	before := time.Now()
	sort.Ints(ints)
	after := time.Now()
	fmt.Println("Sorting takes", after.Sub(before))
	fmt.Println("Is sorted:", sort.IntsAreSorted(ints))
}

func main() {
	// sorting()
	sortLargeArray()
}
