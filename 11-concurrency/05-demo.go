package main

import "fmt"

func printHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
	}
}

func printWorld() {
	for i := 0; i < 5; i++ {
		fmt.Println("World")
	}
}

func main() {
	printHello()
	printWorld()
	fmt.Println("Done")
}
