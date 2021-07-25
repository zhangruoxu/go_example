package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, n := range nums {
		sum += n
	}

	fmt.Println("Sum:", sum)

	m := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range m {
		fmt.Printf("Key %v -> value %v\n", k, v)
	}

	for k := range m {
		fmt.Printf("Key %v\n", k)
	}
}
