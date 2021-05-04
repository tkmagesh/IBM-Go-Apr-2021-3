package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	tick := ticker.C
	stop := time.After(20 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Print("Tick")
		case <-stop:
			fmt.Print("Done")
			return
		default:
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
