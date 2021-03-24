# go

Basic types:
  - `String types`
  - `Numeric types`
  - `Boolean types`

Composite types:
- Aggregation types:
  - `arrays`
  - `structs`
- Reference types:
  - `slices`
  - `maps`
  - `channels`
  - `pointers`
  - `functions`
- Abstract type:
  - `interface`



Value types:
- `int`
- `float`
- `string`
- `bool`
- `structs`
- `array`


## Basic types:

There are the following basic types in Go:
- String types: set of string values.
- Numeric types: integers, floats, bytes and runes.
- Boolean types


#### String types:

`string`: a read-only slice of immutable bytes. Strings are utf-8 encoded by default.

In the context of strings, a `rune` is a unicode code point.

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



## Composite types:


Summary:

**Aggregation types**:
`arrays`: collection of elements that are indexable.
`structs`: groups different (related) types

**Reference types**:
`slices`: collection of elements that are indexable. Has a dynamic length.
`maps`: collection of indexable key-value pairs.
`channels`:
`pointers`: stores the memory address of a value
`functions`: reusabe and organized block of code that performs an action

**Abstract types**:
`interface`: collection of method signatures

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

A slice is actually a slice header/structure that contains 3 fields:
- pointer to a backing array
- length
- capacity

The slice used in go is defined [here](https://github.com/golang/go/blob/master/src/runtime/slice.go).

Note that the slice header contains a pointer. When a slice is passed to a function, the slice header is copied and passed to the function. The backing array however is not copied. 

![Go slice code](https://github.com/saidvandeklundert/go/blob/main/img/slice.png)

These fields describe it's backing array. 
If a slice is copied, the new value will be allocated a new pointer.

When `append()` is used and the new elements do not fit in the curent backing array, Golang allocates a new backing array and copies current content over.

Example slice:

```go
// An empty slice:
var nums []int	
// An array:
var nums [5]int

// The difference: the array has the length defined
```

Very similar to arrays, but notice that the length is not defined at compile time. The lenght of the slice is not a part of it's type.

Slices can only contain one type of element.

```go
// slice literal:
letters := []string{"a", "b", "c", "d"}
```


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

Since methods can be defined on types, a struct can 'equiped' with methods. To add a method to a struct, you need to create a function that takes the struct as a receiver. This is specified between the `func` keyword and the method name.


```go
// Define struct / blueprint
type NetworkDevice struct {
	Name            string
	OperationSystem string
	OsVersion       string
}
// Declare a variable of the type NetworkDevice and set to 0 value:
var R1 NetworkDevice
// Define instance of a struct literal:
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
```
More examples on working with structs [here](https://github.com/saidvandeklundert/go/blob/main/struct.md).

#### functions:

A composite reference data type in go.

Go is a pass by value language. Non-reference types will be copied in memory when the are passed as a value to a function. This applies to receiver values as well as arguments.

When these values are copied, they exist only in the scope of that function. In case you need to change the state of a non-reference value outside of the scope of a function, you will need to pass the pointer to that value.

Functions are first-class citizens in Go.

A function can be passed as an argument to another function.

A method belongs to a type, a function belongs to a package. A method is defined as a function with a reciever argument. The receiver argument precedes the function name:
```go
func (r receiverType) name(s string, i int)(returnString string, returnInt int){
	// We can now work with the receiver in this function
	// the receiver will be of type receiverType
	// we can work with the copy of that type like so: r.attribute
}
```

In effect, this will allow you to run the function/method on the reciever type/object. 

In Python, something similar would be creating a class and then running a method on an instantiated object of that class using `self`:
```python
>>> class MyClass:
...     def name(self):
...         return 'hello world'
... 
>>> 
>>> instanceOfMyClass = MyClass()
>>> instanceOfMyClass.name()
'hello world'
```

Note that this would be similar, it is NOT the same.

```go
func name() {
	// function body, the code goes here
	fmt.Printf("Running the function.")
}

//function with input and output:
func multiply(a, b int)(c int) {	
	c = a * b
	return c
}

// since c is a named result value, it is returned implicitly: 
func multiply(a, b int)(c int) {	
	c = a * b
	return					// aka naked return: returns the named return values
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

Pointers store the memory address of a value. There is a pointer type for every type.

Go is `pass by value`. When a non-reference type is passed to a function, the function will not operate on that data. Instead, it will create a copy and use that copy inside the body of the function. This is giving a certain amount of isolation and immutability. 

It is possible to pass a pointer to a function.This can be used in case you want to have a function change the value of a non-reference type outside the function.

`*` is known as the dereferencing operator. It is used to declare a pointer variable and to access the value stored in the address.

`&` is known as the address operator. It generates a pointer to it's operand.

```go
// the '*' (indirection operator) returns the value of the pointer
// the '&' (address operator) returns the pointer to a value


// Declare integer-type pointer:
var PointerInteger *int

// Declare integer, integer-pointer and then print the values:
var integer int = 100
integerPointer := &integer		// & is used to assign the address of the integer to the pointer
fmt.Printf(`
Pointer memory addres:         %v
Address of the pointer itself: %v
Pointer value:                 %v
`, integerPointer, &integerPointer, *integerPointer)
}
/* ^ gives the following:
Pointer memory addres:         0xc0000ac058
Address of the pointer itself: 0xc0000d8020
Pointer value:                 100
*/


/*
- change someString by changing the pointer value
- assign value to 'value' and 'new_value' by referencing the pointer
*/  
var someString string = "word"		// declare string-literal
pointer := &someString				// declare string-pointer
fmt.Printf("%v\n", pointer) 		// print var value:		 	0xc000088230
fmt.Printf("%T\n", pointer) 		// print variable type:		*string
value := *pointer					//
fmt.Printf("%v\n", value) 			// print var value:			word
fmt.Printf("%T\n", value) 			// print var type:			string  
*pointer = "WORD"					// assign the value to memory address the pointer points to
new_value := *pointer				// short-declare new_value to pointer value 
fmt.Printf("%v\n", pointer)			// print var value:			0xc000088230
fmt.Printf("%v\n", new_value)		// print var value:			WORD  
fmt.Printf("%v\n", someString)		// print var value:			WORD   


```
## Abstract types

All types are concrete types, except the interface type. THe interface type is an abstract type. Unlike with a concrete type, we cannot directly use an interface type to create a value.
#### interfaces:

Specifies a set of 1 or more method signatures. The interface is an abstract type, meaning you cannot create an instance of the interface. 

You can create a variable that is an interface type and that has the methods that belong to the interface. This makes the method a custom type as well as a collection of methods.

Interfaces can be usefull in case a collection of methods do something that can be re-used for a variety of objects.

An example could be in a game where there is an movement method for a knight. Now imagine here are 200+ other characters that also need to move left and right. Instead of copy-pasting the methods over and over again, we can define an interface with a set of methods that can be used by all characters

Example steps to creating and using an interface:
- define the structs that we want to allow use of the interface
- define the functions for the structs
- define the interface and mention the functions that are part of the interface
- define the interface function that references the functions


Interfaces are implicit. We do not manually define the correlation between the interface custom type and the other methods and functions.

Interfaces can be seen as a contract to help manage types. The Go compiler will check all the values and the returns involved, but (obviously) will not check if the logic that is used is sound.


Fun facts:
- if the struct passed as a receiver is not used by the method, we can just pass the type
- structs do not need to have any fields defined. A field-less struct can be used as placeholder for methods and interfaces

It is possible to attach a method to almost any type (even functions!):
- can use pointer and receiver values:
  - int
  - array
  - string
  - struct
  - float64
- use receiver values:
  - slice
  - map
  - func
  - chan

'The bigger the interface, the weaker the abstraction' - Rob Pike.
## Go syntax

### Variables

Typing is everything in Go. Imagine the following:
```
00001010
```

What the above byte means in a Go program can only be understood if the type has been declared.

```go
// Delare variables and set them to their 0 value:
var a string
var b int
var c float64
var d bool

// Declare and initialize variables with short declaration operator:
aa := "string"
bb := 3
cc := 3.14
dd := false
// short declaration operator sometimes referred to as 'productivity operator'
// Does inference. Productivity operator can be used inside a function only

// assign a pre-declared variable
a = "string"

// defining a literal means defining a var together with all the values:
var someString string = "word"

// variable conversion
n := 44
f := float64(n)

// anonymous, or nameless, definition:
```

## Operators

### Comparison operators:

Logical operators combine bool expressions and yield a bool value.

```go
// Logical AND operator:
true && true				// true: when all operands are true

// Logical OR operator:
false || false				// false: one must be True for the expression to be True

// Logical NOT operator:
// If the expression is True, it returns False. If the expression is False, it returns True:
!false						// true
!!false						// false
```


```go
true == true		// true
true != true		// false
```

If a value with a type is not assignable to a variable, then that value cannot be compared with it either.

```go
integer, float := 100, 10.0
integer > float				// not possible!
integer > int(float)		// true, possible since the float is converted
```
### Ordering operators:

```go
1 < 2		// true
2 > 3		// false
2 >= 2		// true
3 <= 2		// false
```

### Comparison operators:

```go
==	// equal to
!=	// not equal to
<	// less than
<=	// less than or equal to
>	// greater than
>=	// greater than or equal to
```

Note: all values of basic types are comparable if they are of the same type.

### Bitwise binary operators:

```go
&	// bitwise AND
|	// bitwise OR
^	// bitwise XOR
&^	// bit clear (AND NOT)
<<	// left shift
>>	// right shift
```
## Flow control

Options:
- if statement
- switch statement
- loop statement

### if

Condition expressions can only be bool

```go

if true {...}	

if true {	
} else if true {
} else if true {
} else if true {
} else if true {
} else {	
}
// Example:
var n int = 1000
if n < 100 {
	fmt.Printf("n is < 100")
} else if n < 200 {
	fmt.Printf("n is < 200")
} else if n < 300 {
	fmt.Printf("n is < 300")
} else {
	fmt.Printf("n >= 300")
}

// Test for multiple conditions:
if word == "word" || word == "WORD"{
	fmt.Printf("word or WORD")
}
```

Simple statement, aka short statement, can be used with an if statement.

```go
// 'Normal':
n, err := strconv.Atoi(os.Args[1])
if err != nil {
	//Handle the error
	fmt.Println("Error", err)
	return
}
fmt.Println("We traversed the happy path! Converted number:", n)
// Simple statement:
if n, err := strconv.Atoi(os.Args[1]); err == nil {
	fmt.Println("We traversed the happy path! Converted number:", n)
}
// notice the separator ';', this syntax is required.
// 'err == nil' > condition expression: so if the error is nill, execute what is in the body
```

Important consideration:

- Variables declared in the simple statement are ONLY available inside the if statement (and it's branches).
- beware of shadowing. Using `:=` you might unintentionaly re-declare a new variable instead of assigning a value to an existing one. Use `=` for the latter.

Shadowing can be caught in VCcode. Add the following to the settings JSON:
```
    "go.lintOnSave": "package",
    "go.vetOnSave": "package",
    "go.vetFlags": [
        "-all",
        "-shadow"
    ]
```


### switch


```go
var word = "word"
switch word {
case "word":
	fmt.Println("word")
case "WORD":
	fmt.Println("WORD")
default:					// location of the default clause does not matter
	fmt.Println("no match")
}
```
The case is like an `if` and the default is like an `else`. The first match is returned, so the case ordering matters. The exception is the location of the default clause. 

You can test for multiple conditions:
```go
case "WORD", "WoRd":			// checked from left to right
	fmt.Println("WORD or WoRd")
```
The values used for case conditions should be unique.

Go automatically converts switch conditions to `if` statements before running.

If constants are used, go will not convert anything. The lookup will be faster.

```go
i := 9
switch {
case i > 100:
	fmt.Println("Bigger then 100")
case i < 10:
	fmt.Println("Less then 10")
	fallthrough						// ignore next case and execute that block also
case i < 100:
	fmt.Println("Less then 100")
case i == 100:
	fmt.Println("100")
default:
	fmt.Println("no match")
}
```

Short switch:

```go
switch i := 9; true {		// short assignment of i and the switch condition (true is default)
case i > 100:
	fmt.Println("Bigger then 100")
default:
	fmt.Println("no match")
}
```

### loop



Basic for loop:
```go
for i := 1; i <= 3; i++ {
	fmt.Printf("word %v\n", i)
}
```
init statement: `i := 1`, simple statement only run once
condition: `i <= 3`, bool expression. The loop continues as long as this remains true.
post statement: `i++`, run after each step of the loop

Note: `i++` is the `IncDec` statement, increments the operand by the untyped constant 1.

The result of the previous loop:
```
word 1
word 2
word 3
```

```go
// basic loop:
for i := 1; i <= 3; i++ {
	fmt.Printf("word %v\n", i)
}
// if i is already defined, it is not needed:
for ; i <= 3; i++ {
	fmt.Printf("word %v\n", i)
}
// IncDec can be moved into the loop:
for ; i <= 3; {
	fmt.Printf("word %v\n", i)
	i++
}
// For statement with single condition:
for i <= 3 {
	fmt.Printf("word %v\n", i)
	i++
}
// infinite loop:
for {
	fmt.Printf("word %v\n", i)
	i++
}
// infinite loop:
var i int
i = 1
for {
	if i > 3 {
		fmt.Printf("Give me a break\n")
		break // break a loop
	}
	fmt.Printf("word %v\n", i)
	i++
}
/* ^ gives:
word 1
word 2
word 3
Give me a break
*/

// use continue to skip an iteration:
if i%2 != 0 {
	i++			
	continue 
}
// Note 'i++', if the check is against i', we should not forget to increment it.
```

## Error handling

`nil`: predeclared identifier. Could be read as 'not yet initialized'. It is like Python's `None`.

All initialized values get a 0-value and for pointer-based types, the 0-value is `nil`.

Typically, a function will return it's usual output and an error. If there is no error, the error value will be `nil`.

Example:
```go

func main() {
	// Convert arg to number:
	n, err := strconv.Atoi(os.Args[1])
	fmt.Println("Converted number:", n)
	fmt.Println("Error", err)
}
```
This will not fail: `go run . 100`:
```
Converted number: 100
Error <nil>
```

We get a response and the error value is `nil`

This waill fail, `go run . a`:
```
Converted number: 0
Error strconv.Atoi: parsing "a": invalid syntax
PS 
```
We get a zero-value returned and the error is not `nil`. We should handle this error and not use the rest of the returned result.

The typical pattern would be:
```go
n, err := strconv.Atoi(os.Args[1])
if err != nil {
	//Handle the error
	fmt.Println("Error", err)
	return
}
fmt.Println("We traversed the happy path! Converted number:", n)
```

## Constants

Constants belong to compile time. 

Constant values are immutable, you cannot change the value of a constant in runtime.

Constants are declared like variables, but with the const keyword you cannot use the `:=` syntax. You must initialize a constant to a value.

Constants can be character, string, boolean, or numeric values.

Untyped constants take the type from their context.

```go
// untyped constants / aka typeless constant
const u_integer = 123
const u_float = 1.2

// typed constants
const t_integer int = 123
const t_float float64 = 1.2

// expressions can be used when initializing constants:
const second = 1
const minute = 60 *s
```

## Stack and heap

[Understanding Allocations: the Stack and the Heap - GopherCon SG 2019](https://www.youtube.com/watch?v=ZMZpH4yT7M0)

Notes from C (need to find good Go resources):

heap:			variable in size, can be used as per need of developer
stack:			function calls and local variables
static/global: 	global variables
code/text: 		stores instructions

stack, static and code memory does not grow.

Stack is divided into frames. Function calls are alloted their own frame.

If more memory then what is available to the stack is required, the program crashes in a stack overflow.

Memory in the heap is controlled by the developper. It can be requested and released.

##

"Make it correct, make it clear, make it concise, make it fast. In that order" - Wes Dyer


# Golang repo

Golang repo is found [here](https://github.com/golang/go).


Go comes with a standard library:
- [documentation](https://golang.org/pkg/)
- [source](https://github.com/golang/go/tree/master/src)



# Go to source and play from your editor

Additional information on a function will appear when you hoover over it. Clicking on the info will offer the option to follow the link to the docs:
![Go to source and play](https://github.com/saidvandeklundert/go/blob/main/img/go_to_source_and_play_1.png)

The docs offer a brief description of the function. Clicking on `Example` will reveal additional information:
![Go to source and play](https://github.com/saidvandeklundert/go/blob/main/img/go_to_source_and_play_2.png)

The full information on the `Println` func looks like this:
![Go to source and play](https://github.com/saidvandeklundert/go/blob/main/img/go_to_source_and_play_3.png)

There are three links here. First is the link to the source code, by clicking on the func name:
![Go to source and play](https://github.com/saidvandeklundert/go/blob/main/img/go_to_source_and_play_4.png)

Additionally we have `Share` and `Run`. Run will run the code in the browser and share will take you to the [Go Playground](https://play.golang.org/):
![Go to source and play](https://github.com/saidvandeklundert/go/blob/main/img/go_to_source_and_play_5.png)