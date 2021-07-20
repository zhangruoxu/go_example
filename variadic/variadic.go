package main

import (
	"fmt"
)

func main() {
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println("Sum:", total)
}