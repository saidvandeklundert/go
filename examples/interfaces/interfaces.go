package main

import "fmt"

//
// 1 define example structs that will all implement the same behavior
// 2 define the methods
// 3 define the interface
// 4 define instances of the structs
// 5 run the method

// 1
type Person struct {
	Name string
}

type ConfidentPerson struct {
	Person
}

type LoudPerson struct {
	Person
}

type SoftSpokenPerson struct {
	Person
}

// 2

func (cp ConfidentPerson) Speaker(s string) {
	fmt.Println("confident voice: ", s)
}

func (lp LoudPerson) Speaker(s string) {
	fmt.Println("loud voice: ", s)
}

func (ssp SoftSpokenPerson) Speaker(s string) {
	fmt.Println("soft voice: ", s)
}

// 3
type Speaker interface {
	Speaker(s string)
}

func main() {
	confidentDave := ConfidentPerson{
		Person{
			Name: "Dave",
		},
	}
	silentBob := SoftSpokenPerson{
		Person: Person{
			Name: "Bob",
		},
	}
	loudLarry := LoudPerson{
		Person{
			Name: "Larry",
		},
	}

	// using the interface:
	confidentDave.Speaker("charge")
	silentBob.Speaker("hi")
	loudLarry.Speaker("HELLO!")

}
