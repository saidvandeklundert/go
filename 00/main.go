package main // package clause should always be the first code
import (
	"fmt" // import fmt
	f "fmt"

	"github.com/saidvandeklundert/go/printstuff" // import fmt as f > f.Print
)

/*

Multiline commenting.

*/
func main() {
	fmt.Println("Said van de Klundert trying on some go.")
	fmt.Println("Gogogogogo, it's a go!", "gogogo")
	fmt.Println("Gogogogogo, it's a go!")
	f.Println("Gogogogogo, it's a go!")
	yolo()
	printstuff.PritStuff()

}
