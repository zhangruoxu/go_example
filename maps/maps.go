package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3

	fmt.Println(m)

	v1 := m["k1"]

	fmt.Println("v1:", v1)

	fmt.Println("Length:", len(m))

	delete(m, "k1")
	fmt.Println(m)

	const K = "k1"
	// Determine whether the map contains the key specified before using the value.
	if value, contains := m[K]; contains {
		fmt.Printf("Key %v value %v\n", K, value)
	} else {
		fmt.Printf("Key %v is not in the map.\n", K)
	}

	// Anther way to create a map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
