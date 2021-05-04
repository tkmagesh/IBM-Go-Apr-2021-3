package main

import (
	"fmt"
	"time"
)

func fibonacci(ch chan int, quit <-chan time.Time) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			time.Sleep(1 * time.Second)
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quitting")
			return
		}
	}

}

func main() {
	ch := make(chan int)
	quit := time.After(15 * time.Second)
	go fibonacci(ch, quit)
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	fmt.Println("Press ENTER to stop")
	fmt.Scanln()

}
