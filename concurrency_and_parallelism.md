### Goroutine:

The Go runtime scheduler manages go routines. Go does not rely on the OS to manage threads. 

Go using it's own scheduler has some advantages:
- Goroutines can be created faster then an OS can create threads
- Goroutines require less memmory. They have a very small initial stack size that can grow as needed.
- switching between Goroutines is faster then swtiching between threads
- the Goroutine scheduler is able to optimize its decisions because it is part of the Go process


A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space. Additionally, context switching between go routines is also a lot more lightweight versus having the OS context switch threads. More [here](https://golang.org/doc/effective_go#goroutines).

W. Kennedy's take: Go has turned IO-bound work into CPU-bound work because by running multiple Go routines in a single thread (from an OS pov), Go can keep the CPU busy all the time.

W. Kennedy's take on 5 kids (from Go scheduler mechanics): "I just travel a lot, I get out of town."
Goroutine vs threads:

| Go routine  | thread  |
|---|---|
| Managed by the Go runtime   | Managed by kernal  |
| Not hardware dependant   | hardware dependant  |
| Easy communication through channels  | No easy communication between threads  |
| Low latcency communication between channels | Latency involved in inter-thread communication  |
| Go routine does not have an ID as it does not have thread local storage  | unique ID because they have thread local storage  |
| Cheap | Expensive (less then processes, but still)  |
| fast startup  | slow startup  |
| growable segmented stacks | no growable segmented stacks |
| cooperatively scheduled  | preemptively scheduled |




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




Go scheduler is in charge of scheduling Go routines.

Concurrency is not parallelism.

Concurrency: multiple threads executing code.

Parallelism: multiple threads executed at the exact same time (requires multiple CPUs).

Important note: channels are blocking.

Go concurrency slogan:

```
Do not communicate by sharing memory; instead, share memory by communicating.
```
