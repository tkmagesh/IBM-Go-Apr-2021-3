package main

import (
	"fmt"
)

/*
	"Communicate by sharing memory"
*/

//simplified version of 03-demo.go

/*
	By default
		reading and writing data in a channel (non buffered) is blocking operation
		write operation can happen ONLY if a read operation is initated (non buffered channel)
		read operation can succeed ONLY if the write operation is completed
*/

func add(x, y int, ch chan int) {
	result := x + y
	//writing data into the channel
	ch <- result
}

func main() {
	ch := make(chan int)
	fmt.Println("Triggering the add fn")
	go add(100, 200, ch)
	//reading data from the channel
	result := <-ch
	fmt.Println("Result = ", result)
}
