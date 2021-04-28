First major programming language ever built to take advantage of multiple cores. Concurrency is at the heart of the design of the Go programming language.

Go enables two styles of concurrent programming. CSP, or communicating sequential processes, is a model of concurrency in which values are passed between independant activities (goroutines) but variavles are for the most part confined to a single activity. Go routines and channels follow this CSP model.

Go scheduler is in charge of scheduling Go routines.

Concurrency is not parallelism.

Concurrency: multiple threads executing code.

Parallelism: multiple threads executed at the exact same time (requires multiple CPUs).

Important note: channels are blocking.

Go concurrency slogan:

```
Do not communicate by sharing memory; instead, share memory by communicating.
```

In the context

`Goroutine`:

A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space. More [here](https://golang.org/doc/effective_go#goroutines).

A `go` statement starts the execution of a function call as an independant concurrent thread of control, or `goroutine`, within the same address space.


Example staring a go routine and using wait group to ensure everything finishes:
```go
// Start 6 Go routines that wait 5 seconds. This function completes in less then 5 seconds:
func ConcurrentExample() {
	// define the wait group:
	var wg sync.WaitGroup
	fmt.Printf("\nStart ConcurrentExample Func")
	// Start the Go routine and pass the address of the wait group:
	go ConcurrentFunc(1, &wg)
	go ConcurrentFunc(2, &wg)
	go ConcurrentFunc(3, &wg)
	go ConcurrentFunc(4, &wg)
	go ConcurrentFunc(5, &wg)
	go ConcurrentFunc(6, &wg)
	// specify the number of wait groups:
	wg.Add(6)
	// Wait for all of them to finish:
	wg.Wait()
	fmt.Printf("\nFinished ConcurrentExample Func")
}

// Print a number and wait 5 seconds:
func ConcurrentFunc(Nr int, wg *sync.WaitGroup) {
	// Ensure we wrap up saying done even if we crash
	defer wg.Done()
	fmt.Printf("\nStart ConcurrentFunc %d", Nr)
	time.Sleep(5 * time.Second)
	fmt.Printf("\nFinished ConcurrentFunc %d", Nr)
}
```
This example is found [here](https://github.com/saidvandeklundert/go/blob/main/examples/wait_group_simple.go). A more elaborate example is found [here](https://github.com/saidvandeklundert/go/blob/main/examples/wait_group.go).

Look into Mutex or Atomic to avoid race conditions in case the Go routines share variables.


Notes:

You can use the comma idiom to check if a channel is empty:
`v, ok := <-ch`

`ok` is `false` if there are no more values to receive and the channel is closed.

Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

You can go from general to specifics with channels:
```go
c := make(chan int) // bidirectional channel is created


go send(c)		// bidirectional channel is passed in

func send(c chan<- int)	// send channel
```

Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

Context can be used to ensure spawned Go routines are halted. This to prevent leaking of Go routines and to prevent wasting of resources.


## Channels

Channels offer a way to manage concurrent code. They allow us to pass values between goroutines.

Channels are blocking. This means that after you push a message onto a channel, the operation is blocking untill it is read.



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