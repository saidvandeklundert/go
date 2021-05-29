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

func (cp ConfidentPerson) Speak(s string) {
	fmt.Println("Speaks with a confident voice: ", s)
}

func (lp LoudPerson) Speak(s string) {
	fmt.Println("Speaks with a loud voice: ", s)
}

func (sp SoftSpokenPerson) Speak(s string) {
	fmt.Println("Speaks with a soft voice: ", s)
}

func (sp SoftSpokenPerson) Greet(s string) (g string) {
	g = fmt.Sprintf("Greets with a soft voice: %s", s)
	return g
}

func (sp SoftSpokenPerson) Sleep() {
	fmt.Println("zzzzzzzzzzzzzzzzzzzz")
}

// 3
type HumanBehaviors interface {
	Speak(s string)
	Greet(s string) (g string)
	Sleep()
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
	confidentDave.Speak("charge")
	silentBob.Speak("talk talk talk")
	g := silentBob.Greet("Hello there.")
	fmt.Println(g)
	silentBob.Sleep()
	loudLarry.Speak("WOW THIS IS AMAZING!")

}
