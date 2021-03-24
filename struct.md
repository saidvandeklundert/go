A struct is a composite aggregate type.


```go
// Define struct
type Human struct {
	Name       string
	Surname    string
	Age        int
	HairColour string
	Children   []string
}
// Declare a variable of the type NetworkDevice and set to 0 value:
var Marie Human
// Define instance of a struct literal:
Marie := Human{
	Name:            "Marie",
	Surname: "van de Klundert",
	Age:       2,
}

```