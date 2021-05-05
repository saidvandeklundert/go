Functions are first-class citizens in Go.





```go
// Simple function:

//function with input and output:
func multiply(a, b int)(c int) {	
	c = a * b
	return c
}


// Closure: 
//  The inner function uses the scope of the outer function.
//  The scope will persist as long as the function persists.
func CreateCounter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}
// create the counter by assigning the
// anonymous function to the counter variable:
counter := CreateCounter()
counter()           // 1
counter()           // 2
counter()           // 3


// Recursion example:
func RecEx(a, b int) (c int) {
	c = a * b
	if c < 100 {
		return RecEx(c, b)
	} else {
		return c
	}
}

RecEx(2, 2)         // 128: returns RecEx 5 x before it is > then 100, then returns c
RecEx(2, 100)       // 200: returns c immediately as it is > then 100



// Function literal:
func() {
	fmt.Println("inside the unnamed func")
}

// EXECUTING a function literal:
func() {
	fmt.Println("inside the unnamed func")
}() // <-()
```