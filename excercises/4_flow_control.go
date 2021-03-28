package main

import "fmt"

func main() {
	// 1:
	// Define a slice containing a, b, c and d.
	// Loop over the slice and print both the index as well as the value:
	var Slice = []string{"a", "b", "c", "d"}
	for index, value := range Slice {
		fmt.Println(index, value)
	}
	// 2:
	//
	// Loop over the following array
	// 1, 2, 3, 4, 5, 6
	//
	// Iterate the array and print every value (not the index)
	// When the value is 3, skip it.
	// Be sure to exit the loop before the value reaches 6
	//
	Array := [6]int{1, 2, 3, 4, 5, 6}
	for _, nr := range Array {
		if nr == 3 {
			continue
		} else if nr >= 6 {
			break
		}
		fmt.Println(nr)
	}

}
