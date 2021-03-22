package main

import "fmt"

func main() {
	fmt.Println("Some example functions.")
	var str_value string = "word"
	echo(str_value)
	var int_value int = 100
	new_int := multiply(int_value)
	// Explaining what multiply returned:
	fmt.Printf(`
int_value now:                 %v
int_ptr:                       %v
new_int:                       %v
new_int pointer value:		   %v

`, int_value, &int_value, new_int, &new_int)
	// Old int left as is
	// Change int_value:
	multiplyExisting(&int_value)
	fmt.Printf(`
int_value now:                 %v
int_ptr:                       %v

`, int_value, &int_value)

	// Illustrate we can pass a func + arg as arg to a func:
	x_1 := returnInt(silent_multiply(int_value))
	fmt.Printf("\nreturnInt(silent_multiply(int_value)) got us: %d\n", x_1)

	// Illustrate naked return:
	_ = echoReturnString(str_value)

}

// Echo a string:
func echo(s string) {
	fmt.Printf("%s\n", s)
}

// Return an integer multiplied by 2:
func multiply(i int) int {
	i = i * 2
	fmt.Printf("multiply i value: %d", i)
	return i
}

// multiply the int passed to the func:
func multiplyExisting(i *int) {
	*i = *i * 2
	fmt.Printf("multiplyExisting i value %d", &i)
}

// Return an int:
func returnInt(i int) int {
	return i
}

func silent_multiply(i int) int {
	i = i * 2
	return i
}

// Naked return:
func echoReturnString(s string) (str string) {
	str = s
	fmt.Printf("%s\n", str)
	return
	// Due to (str string) this is the same as return str

}
