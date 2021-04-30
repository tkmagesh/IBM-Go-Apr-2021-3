package main

import "fmt"

func main() {
	greet("Magesh", "Have a good day!")
	fmt.Printf("Adding 10 and 20 results in %d\n", add(10, 20))
	quotient, remainder := divide(80, 9)
	fmt.Printf("Dividing 80 by 9, quotient=%d, remainder=%d\n", quotient, remainder)
}

/* func greet(name string, message string) { */
func greet(name, message string) {
	fmt.Printf("Hi %s, %s\n", name, message)
}

/*
func add(x, y int) int {
	return x + y
}
*/

func add(x, y int) (result int) {
	result = x + y
	return
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
