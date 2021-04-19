### Error handling in Go

Go does not have exceptions, as stated [here](https://golang.org/doc/faq#exceptions):

```
We believe that coupling exceptions to a control structure, as in the try-catch-finally idiom, results in convoluted code. It also tends to encourage programmers to label too many ordinary errors, such as failing to open a file, as exceptional.
```

Error is a type in Go, it is decribed [here](https://pkg.go.dev/builtin#error):

`type error`:

```go
type error interface {
	Error() string
}
```

This can be used to create your own custom errors. 

The `errors` packages is part of the standard library and the documentation is found [here](https://pkg.go.dev/errors).

The example listed in the package is something like this:
```go
import "errors"

err := errors.New("emit macho dwarf: elf header corrupted")
```

A nice thing to know is what is happening 'under the hood'. The following is the source code:
```go
// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
	return &errorString{text}
}
// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```


Checking for an error and logging fatal:
```go
import "log"
if err != nil {
	log.Fatal(err)
	log.Printf("%v", err)
}
```

Instead of logging an error, you can also inspect the error and initiate followup actions:
```go
if err != nil {
	if strings.Contains(err.Error(), "404") {
		Error.Println(err)
```