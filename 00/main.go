package main // package clause should always be the first code
import (
	"fmt" // import fmt
	f "fmt"
)

/*

Multiline commenting.

*/

var s string

func main() {
	s = "Said van de Klundert trying on some go."
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println("Gogogogogo, it's a go!", "gogogo")
	f.Println("Gogogogogo, it's a go!")
	yolo()
	PrintStuff()
	fmt.Println(version())
}
