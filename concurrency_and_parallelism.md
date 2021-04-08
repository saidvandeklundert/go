First major programming language ever built to take advantage of multiple cores. COncurrency is at the heart of the design of the Go programming language.

Go scheduler is in charge of scheduling Go routines.

Concurrency is not parallelism.

Concurrency: multiple threads executing code.

Parallelism: multiple threads executed at the exact same time (requires multiple CPUs).

Important note: channels are blocking.

Go concurrency slogan:

```
Do not communicate by sharing memory; instead, share memory by communicating.
```

`Goroutine`:

A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space. More [here](https://golang.org/doc/effective_go#goroutines).

A `go` statement starts the execution of a function call as an independant concurrent thread of control, or `goroutine`, within the same address space.

Example staring a go routine and using wait group to ensure everything finishes:
```go
```


Example using channels:
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	/* create a slice */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	c := make(chan int)

	for _, number := range numbers {
		go channlFunc(number, c)
	}

	for range numbers {
		fmt.Println(<-c)
	}
}

func channlFunc(number int, c chan int) {
	time.Sleep(2 * time.Second)
	c <- number
}

```

Receive channel:
- can receive values from the channel
- cannot close a recieve channel

Send channel:
- push values to the channel
- cannot receive/pull or read from the channel