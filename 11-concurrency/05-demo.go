package main

import (
	"fmt"
	"time"
)

func print(msg string, in chan string, out chan string) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(1 * time.Second)
		fmt.Println(msg)
		out <- "done"
	}
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	go print("Hello", ch1, ch2)
	go print("World", ch2, ch1)
	ch1 <- "start"
	fmt.Println("Press ENTER...")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}
