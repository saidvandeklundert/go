package main

import (
	"fmt"
	"runtime"
)

func yolo() {
	fmt.Println("Yolo! Scoping thing.")
	fmt.Println("I have this many CPU:")
	fmt.Println(runtime.NumCPU())
}
