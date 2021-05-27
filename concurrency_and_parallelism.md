## Introduction to concurrency in Go

Go's main concurrency model is based on CSP, Communication Seqeuntial Processes.
Concurrency is not parallelism:
- concurrency: multiple threads executing code.
- parallelism: multiple threads executed at the exact same time (requires multiple CPUs).

In the context of Go, the following terms are important:
- `Goroutines`: User-space thread managed by the Go runtime. The thread is not managed by the OS. Goroutines are used to execute tasks independently.
- `Channels`: reference type that is used for communication and synchronization between Goroutines.
- `Go scheduler`: multiplexes Goroutines onto OS threads.

*Channels*:
- are Goroutine-safe
- store and pass values between Goroutines
- provide FIFO semantics (when they are buffered)
- can cause Goroutines to block and unblock

The zero value for a channel is `nil`.


### Goroutines

A Goroutine can be seen as a lightweight process managed by the Go runtime. Scheduling of goroutines onto OS threads is handled by the Go runtime scheduler. The runtime scheduler multiplexes Go routines into OS threads. 

The Go runtime scheduler manages go routines. Go does not rely on the OS to manage threads. Go using it's own scheduler has some advantages:
- Goroutines can be created faster then an OS can create threads
- Goroutines require less memory. They have a very small initial stack size that can grow as needed.
- switching between Goroutines is faster then swtiching between threads
- the Goroutine scheduler is able to optimize its decisions because it is part of the Go process


A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space. Additionally, context switching between go routines is also a lot more lightweight versus having the OS context switch threads. More [here](https://golang.org/doc/effective_go#goroutines).

The `go` statement starts the execution of a function call as an independant concurrent thread of control, or `goroutine`, within the same address space.

`W. Kennedy`: Go has turned IO-bound work into CPU-bound work because by running multiple Go routines in a single thread (from an OS pov), Go can keep the CPU busy all the time.

`Goroutine vs OS threads`:

| Go routine  | OS thread  |
|---|---|
| Managed by the Go runtime   | Managed by kernal  |
| Not hardware dependant   | hardware dependant  |
| Easy communication through channels  | No easy communication between threads  |
| Low latency communication between channels | Latency involved in inter-thread communication  |
| Go routine does not have an ID as it does not have thread local storage  | unique ID because they have thread local storage  |
| Cheap | Expensive (less then processes, but still)  |
| fast startup  | slow startup  |
| growable segmented stacks | no growable segmented stacks |
| cooperatively scheduled  | preemptively scheduled |


Look into Mutex or Atomic to avoid race conditions in case the Go routines share variables. 

Channels enable safe communication between Go routines. 

Go also offers mutex (mutual exclusion). Use a mutex when multiple Go routines read/write to a shared value (struct field for instance). Using mutexes, when a Go routine needs to access a value shared between Go routines, it needs to call `Lock` on it. The Goroutine aquires the lock immediately or after waiting for other Goroutines to release it. Once it lock is aquired, the Go routine can safely access and work with the value. When the Go routine is done, it can call `Unlock` the value.

When to use channels and when to use mutexes:
- when sharing a field in a struct, use mutexes
- when channels stand in the way of fixing a critical performance issue, use mutexes
- if you are coordinating goroutines or tracking a value that is transformed by a series of Goroutines, use channels

In addition to mutexes and channels, there is also atomics. According to 'Learning Go', skipp this unless you are a concurrency expert solving a problem where every bit of performance matters.
## Channels

Channels in go facilitate communication between different Go routines. In a sense, channels offer a way to manage concurrent code. They allow us to safely pass values between Goroutines. 

The channel is a reference type. And in addition to that, it is defined for a specific type. The zero value for a channel is `nil`.

Channels are blocking. This means that after you push a message onto a channel, the operation is blocking untill it is read.

After a channel is created, and a go routine is started, you can use the channel to:
- send data into the channel
- recieve values in the channel


`Receive channel`:
- can receive values from the channel
- cannot close a recieve channel

`Send channel`:
- push values to the channel
- cannot receive/pull or read from the channel


You can use the comma idiom to check if a channel is empty:
`v, ok := <-ch`

`ok` is `false` if there are no more values to receive and the channel is closed.

Sending on a closed channel will cause a panic so only the sender should close a channel. Closing is generally only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop. 

Context can be used to ensure spawned Go routines are halted. This to prevent leaking of Go routines and to prevent wasting of resources.

## Examples



You can go from general to specifics with channels:
```go
c := make(chan int) // bidirectional channel is created


go send(c)		// bidirectional channel is passed in

func send(c chan<- int)	// send channel
```

Example on how to make a channel:
```go
c := make(chan string)
```

Example on starting go routines:

```go
for _, link := range links {
	go touchLink(link, c)
}
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


Go concurrency slogan:

```
Do not communicate by sharing memory; instead, share memory by communicating.
```

Inspiriation for these notes came from:
- Learning Go
- Go in action
- `GopherCon 2017: Kavya Joshi - Understanding Channels`
- `GopherCon 2018: Kavya Joshi - The Scheduler Saga`
