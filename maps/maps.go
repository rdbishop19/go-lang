package main

import (
	"fmt"
	"maps"
)

func main() {

	// create an empty map
	m := make(map[string]int)

	// set key/value pairs
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// get a value
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// if key doesn't exist, you'll get the zero value for the value type
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// builtin len returns the number of key/value pairs when called on a map
	fmt.Println("len:", len(m))

	// builtin delete removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// remove all key/value pairs from a map
	clear(m)
	fmt.Println("map:", m)

	// optional second return value when getting a value from a map
	// indicates if the key was present in the map
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// utility function example
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
