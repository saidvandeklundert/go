
## Go CLI:

go run:
go build:
go fmt:
go install:
go get:
go test:



## Packages:

The first line of a file must always declare what package it belongs to.

Package 'main' indicates a package is an executable that can be compiled. The package MUST have a function called main.

Packages with oter names can be used as a de

```go
import "fmt"
import (
	"fmt"
	"sort"
)
```