package main

import "fmt"

func main() {
	adder := getAdder(100)
	fmt.Println(adder(200))
}

func add(x, y int) int {
	return x + y
}

func getAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
