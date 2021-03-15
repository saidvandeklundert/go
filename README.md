# go


### Go types:

There are the following basic types in Go:
- String types: set of string values.
- Numeric types: integers, floats, bytes and runes.
- Boolean types


#### String types:

`string`: a read-only slice of immutable bytes. Strings are utf-8 encoded by default.


#### Numeric types:

Numeric integer types:
`int`: integer literals without a fractional part( -2 0 88 )
`int8`: 8 bit integer
`int16`: 16 bit integer
`int32`: 32 bit integer
`int64`: 64 bit integer

All the above also have an unsigned variant ( `uint`, `uint8`, etc.). Unsigned integers can only contain positive numbers.

Numeric floating types:

Floats:
`float32`: integer literals wit a fractional part (-1.7 0. 88.8)
`float64`: integer literals wit a fractional part (-1.7 0. 88.8)

Complex numbers:
`complex64`: two float32's
`complex128`: two float64's

`byte`: can store 8 bits
`uintptr`: unsigned integer type


`aliased types`: aliased for readability:
 - `byte` == `uint8`
 - `rune` == `int32`


#### Boolean types:

`bool`: pre-declared constant that can be true or false. (1 byte)



### Composite types

What is `string internals`: Byte slices, ASCII & Unicode encoding and decoding.


#### summary:

Aggregation types:
`arrays`: collection of elements that are indexable.
`structs`: groups different (related) types

Reference types:
`slices`: collection of elements that are indexable. Has a dynamic length.
`maps`: collection of indexable key-value pairs.
`channels`:
`pointers`: stores the memory address of a value
`functions`: reusabe and organized block of code that performs an action

Interface:
`?`

#### arrays:

`arrays`: collection of elements that are indexable. 

It has a fixed length and stores elements of the same type. Array properties include the type of elements stored and the lenght of the array. Since the lenght of the array is a property, it belongs to compile time.

When arrays are copied, the value of the array is copied along with it. The 

Arrays populate conitguous memory locations.
```go
// Array with 10 empty string elements
var emptyArrayStr [10]string
// Array with 10 empty integer elements
var emptyArrayInt [10]int
// Array with several strings and several empty string elements
var arrayStrings [10]string{"Some", "values", "stored", "in", "the", "array"}
// Shorthand assignment
arrayStringsShort := [10]string{"Some", "values", "stored", "in", "the", "array"}
//Assign strings to an array and iterate the array:
var arrayStrings = [5]string{"Iterating", "an", "array", "in", "Go"}
for index, word := range arrayStrings {
	fmt.Println(index, word)
}
//Assign the order of the strings inside the array:
var arrayStrings = [5]string{
	0: "Iterating",
	4: "Go",
	2: "array",
	3: "in",
	1: "an",
	}
for index, word := range arrayStrings {
	fmt.Println(index, word)
}
//ellipsis sets the array length based on the items assigned to it
var arrayStrings = [...]string{"Iterating", "an", "array",}
for index, word := range arrayStrings {
	fmt.Println(index, word)
}
fmt.Println(len(elipsisArrayStrings)) // 3

// Arrays can be multidimensional:
servers := [2][2]string{
	[2]string{"container-1", "ubi"},
	[2]string{"container-2", "centos"},
}
// Multidimensional arrays are better defined using the elipsis and declaring the type only once:
servers := [...][2]string{
	{"container-1", "ubi"},
	{"container-2", "centos"},
}
```


#### slices:

`slices`: collection of elements that are indexable. Has a dynamic length.

An array cannot change it's size. A slice can change in runtime.

The slice header contains the pointer. If a slice is copied, the new value will be allocated a new pointer.

Example slice:

```go
var nums []int	// empty slice, no length defined
var nums [5]int	// array, length defined
```

Very similar to arrays, but notice that the length is not defined at compile time. The lenght of the slice is not a part of it's type.

Slices can only contain one type of element


#### maps:

`maps`: collection of indexable key-value pairs. The map is a hash-table. 

They `key` value can be any comparable, or phrased differently, any type that supports the equality operator.

A map variable is a pointer to a map header value in memory. The map header tracks everything related to the map.

When you pass a map value to something (a function for instance), you are only passing the pointer.

```go
// declare map
var dict map[string]string

// declare map using make
var dict = make(map[string]string)

// initialize map	
dict := map[string]string{
	"r1": "juniper",
	"r2": "arista",
	"r3": "cisco",
}

// print map values to screen
fmt.Printf("%v\n", dict)

// print Go-syntax representation to screen
fmt.Printf("%#v\n", dict)

// Safely access key:
if value, ok := dict["r4"]; ok {
	fmt.Printf("The value is %v.\n", value)
}

// Safely access key and return message if it is not found:
if value, ok := dict["r4"]; ok {
	fmt.Printf("The value is %v.\n", value)
} else {
fmt.Println("Value not found.")
}

// Loop over key/values"
for key, value := range dict {
    fmt.Printf("%v %v\n", key, value)
}
// delete a key (safe when referencing non-existing keys)
delete(dict, "r233")

// copy dict
var dict2 = make(map[string]string)
for key, value := range dict {
	dict2[key] = value
}

// remove the pointer from the var (does not destroy the map!)
dict = nil
// destroy the map
for key := range dict {
	delete(dict, key)
}

```


#### structs:

`structs`: groups different (related) types

Allows you to define a static blueprint. 

The field name and field types are defined at compile time. 

The field values belong to runtime and can change during program execution.

Structs may store different types of data. 

Structs can contain structs as well.

In OOP, that is inheritence with an "is a" relationship. In Go, there is no inheritence. In Go, go for composition. Go allows for structs to be embeddeded into other structs. The 1 struct 'has-a' struct embedded.


| Slices and maps            | structs                    |
| -------------------------- |:--------------------------:|
| single element type        | different types of fields  |
| Dynamic number of elements | fixed number of fields     |

```go
// Define struct / blueprint
type NetworkDevice struct {
	Name            string
	OperationSystem string
	OsVersion       string
}
// Define instance of a struct:
R1 := NetworkDevice{
	Name:            "R1",
	OperationSystem: "junos",
	OsVersion:       "16.R4",
}
// print Go-syntax representation of instance to screen
fmt.Printf("struct: %#v\n", R1)

// print individual struct values:
fmt.Println(R1.Name)			// R1
fmt.Println(R1.OperationSystem)	// junos
fmt.Println(R1.OsVersion)		// 16.R4

// Change or assign an individual struct value:
R1.Name = "Router1"
fmt.Println(R1.Name)			// Router1

// No need to assign values to all fields:
R2 := NetworkDevice{
	Name:            "R2",
	OperationSystem: "eos",
}
fmt.Printf("%#v", R2)   // main.NetworkDevice{Name:"R2", OperationSystem:"eos", OsVersion:""}

// Struct values are copied along with their fields:
R3 := R2
R3.Name = "R3" // This changes the name for R3 only

// Structs are equal if all fields are equal
R2 == R3		// false

// unless you specifically compare fields
R2.OperationSystem == R3.OperationSystem  // true

//
type Packages struct {
	language string
	packages []string
}
type Container struct {
	Packages        Packages
	Name            string
	OperationSystem string
	OsVersion       string
}

container := Container{
	Packages:        Packages{language: "python3", packages: []string{"napalm", "fastapi", "netmiko"}},
	Name:            "Golem",
	OperationSystem: "CentOS",
	OsVersion:       "7",
}
fmt.Printf("struct: %#v\n", container) // Gives the following:
/*
struct: main.Container{Packages:main.Packages{language:"python3", packages:[]string{"napalm", "fastapi", "netmiko"}}, Name:"Golem", OperationSystem:"CentOS", OsVersion:"7"} 
*/
//access item in a slice in a nested struct:
fmt.Printf("struct: %#v\n", container.Packages.packages[1])

// Embedding a struct with an anonomous field:
type Container struct {
	Packages        
	Name            string
	OperationSystem string
	OsVersion       string
}
// An anonymous field gets the name from it's type:
//accessing an item in a slice in nested struct defined as an anonymous field:
fmt.Printf("struct: %#v\n", container.Packages.packages[1])
fmt.Printf("struct: %#v\n", container.packages[1])
```

#### functions:

A composite reference data type in go.

Go is a pass by value language. Inputs to a function are local to that function.

```go
func name() {
	// function body, the code goes here
	fmt.Printf("Running the function.")
}

// a, b are integer input arguments, c is declared as integer output:
func multiply(a, b int)(c int) {	
	c = a * b
	return c
}

// since c is a named result value, the return will implicitly return c: 
func multiply(a, b int)(c int) {	
	c = a * b
	return
}


// print the function or store the result in a var:
fmt.Println(multiply(2, 6))
c = multiply(2, 6)

// Mind Go's pass by value, important when dealing with functions that change aggregate values.
// Here is a struct:
type SomeStruct struct {
	Name string
}
// Func that changes struct:
func changeSomeStruct(s SomeStruct) SomeStruct {
	s.Name = "new_name"
	return s
}
// Define struct, print it to screen, change it and print it to screen again:
s := SomeStruct{Name: "name"}
fmt.Printf("struct: %#v\n", s.Name) 	// struct: "name"
s = changeSomeStruct(s)					// pass the struct (s) value and re-assign it to s
fmt.Printf("struct: %#v\n", s.Name)		// struct: "new_name"


// Now the same with a map (map is reference type):
func changeMapping(m map[string]string) {
	delete(m, "a")
	return
}
mapping := map[string]string{"a": "a", "b": "b"}
fmt.Printf("mapping: %#v\n", mapping)		//mapping: map[string]string{"a":"a", "b":"b"}
changeMapping(mapping)
fmt.Printf("mapping: %#v\n", mapping)		// mapping: map[string]string{"b":"b"}

```


#### pointers:

There is a pointer type for every type.

```go
var someString string = "word"
/*
	- create a pointer variable
	- assign the memory address of someBytes to it
	- the '&' (address operator) returns the pointer to a value
*/
pointer := &someString
fmt.Printf("%v\n", pointer) // 0xc000104230 (example memeroy address)
fmt.Printf("%T\n", pointer) //*string
/*
	- provide the value that a memory address refers to
	- the '*' (indirection operator) returns the value of the pointer
*/
value := *pointer
fmt.Printf("%v\n", value) // word
fmt.Printf("%T\n", value) // string
/*
	- update the value directly through the pointer
*/
*pointer = "WORD"
new_value := *pointer
fmt.Printf("%v\n", pointer)
fmt.Printf("%v\n", new_value) // word
	
```