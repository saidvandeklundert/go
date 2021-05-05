package main

import "fmt"

func main() {

	// Create a slice with a length of 5 elements and a capacity of 8.
	fruits := make([]string, 5, 8)

	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	inspectSlice(fruits)

	fruitSlice := fruits[1:3] // Orange, Banana, same backing array
	inspectSlice(fruitSlice)
	NewSlice := make([]string, len(fruits))
	copy(NewSlice, fruits) // copy fruits to NewSlice, NewSlice uses new backing array
	inspectSlice(NewSlice)

}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			s)
	}
}
