package main

import (
	"fmt"
	"sync"
)

/*
	The below is "sharing memory for communication" - not advicable
	instead "Communicate by sharing memory" (03-demo.go)
*/
var result int
var wg *sync.WaitGroup = &sync.WaitGroup{}

func add(x, y int) {
	result = x + y
	wg.Done()
}

func main() {
	wg.Add(1)
	fmt.Println("Triggering the add fn")
	go add(100, 200)
	wg.Wait()
	fmt.Println("Result = ", result)
}
