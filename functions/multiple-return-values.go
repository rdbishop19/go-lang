package main

import "fmt"

// The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// multiple return values with _multiple assignment_
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// use _ to ignore one of the return values
	_, c := vals()
	fmt.Println(c)
}
