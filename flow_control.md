

How the computer reads the program. For flow control, Go offers the following options:
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


Anatomy of a for loop:
```go
// initializer: initializes the loop, e.g. i := 1
// condition: tests for a condition, e.g. i <= 3
// post: example could be the incdec, e.g. i++
for init; condition; post{
    fmt.Println("yolo")
}
```

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
// loop using range:
for index, value := range l {
	fmt.Println(index, value)
}
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

```go
panic("Stop the program right here!")

if err != nil {
    panic(err)
}
```