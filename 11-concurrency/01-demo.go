package main

import (
	"fmt"
)

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	go print("Hello")
	go print("World")
	//time.Sleep(2 * time.Second)
	fmt.Println("press ENTER to continue..")
	var input string
	fmt.Scanln(&input)
}
