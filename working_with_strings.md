

### Contains:

Used to check for presence of substrings:
```go
import "strings"
exampleString := "Some words uttered by someone."
strings.Contains(exampleString, "words")                    // true
strings.Contains(exampleString, "giant massive robot")      // false
```