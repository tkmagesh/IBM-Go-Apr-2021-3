package main

import (
	"fmt"
	"sync"
)

/*
	"Communicate by sharing memory"
*/

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	result := x + y
	//writing data into the channel
	ch <- result
	wg.Done()
}

func main() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	fmt.Println("Triggering the add fn")
	go add(100, 200, wg, ch)
	//reading data from the channel
	result := <-ch
	wg.Wait()
	fmt.Println("Result = ", result)
}
