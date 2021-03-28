package main

func main() {
	// Basic types
	//
	// Define a string, an integer and a float64

	/*fmt.Printf(`
	%T is %v
	%T is %v
	%T is %v
	`,
			String, String, Int, Int, Float, Float)
	*/
	// Short declare a new string and re-assign the value of the previous string:

	//fmt.Printf(`%v is %v`, NewString, String)
	// Reference types:
	// Define a slice containing "a", "b", "c"

	// Define a map that maps a, b, and c to 1,2 and 3.

	/*fmt.Printf(`
	%T is %v
	%T is %v
	`, Slice, Slice, Map, Map)/*
		// Create a pointer to the integer 'intPointer':


		/*fmt.Printf(`
	intPointer is a %T

	The memory address of the pointer is %v
	The memory address of the value is %v
	The value is %v

	`, intPointer, &intPointer, intPointer, *intPointer)*/

	// Change the value of the pointer through the pointer:

	// Assign the new value to a new variable and print it:

	//fmt.Println(NewValue, Int)

	// Composite or aggregation types:
	//
	// Define an empty array with room for 10 string elements:

	//fmt.Println(Array)
	// Put "hello" in the second element and world in the 10th element

	// Create a struct called 'Human'.
	// The fields that are required are the following:
	// Name: string					Jan
	// Age: integer					35
	// Children: slice				Joe, Anne
	// Favorites: mapping			Color: blue, Food: pizza
	// ContactInfo: struct
	//	- email, string				jan@yolo.com
	//	- phone number, string		06121345678
	//

	// Print the type, the value and the Go-syntax for this struct:
	/*fmt.Printf(`
	It is a %T
	The value is %v
	The syntax is %#v
	`, Jan, Jan, Jan)*/

}
