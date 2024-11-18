package main

import (
	"fmt"
	"slices"
)

func main() {

	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Use `make` to create a slice of a specific length and capacity.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Set and get elements just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length (just like python)
	fmt.Println("len:", len(s))

	// `append` returns a new slice with the added element(s).
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can be copied.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax `slice[low:high]`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l)

	// This slices up from (and including) `s[2]`.
	l = s[2:]
	fmt.Println("sl3:", l)

	// We can declare and initialize a slice in one line.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// The slices pacakge contains a numer of useful utility functions for slices.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slices can be composed into multi-dimensional data structures.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
