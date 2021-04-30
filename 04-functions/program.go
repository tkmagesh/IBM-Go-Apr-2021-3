package main

import (
	"fmt"
	"func-demos/calculator"
	gt "func-demos/greeter"
	"func-demos/math"
	"func-demos/utils"
)

func main() {
	gt.Greet("Magesh", "Have a good day!")
	fmt.Println(calculator.PkgName)
	/*
		loggerAdd := logOperation(calculator.Add)
		fmt.Printf("Adding 10 and 20 results in %d\n", loggerAdd(10, 20))
	*/
	fmt.Printf("Adding 10 and 20 results in %d\n", logOperation(calculator.Add)(10, 20))

	loggerSubtract := logOperation(calculator.Subtract)
	fmt.Printf("Subtracting 10 from 20 results in %d\n", loggerSubtract(20, 10))

	quotient, remainder := calculator.Divide(80, 9)
	fmt.Printf("Dividing 80 by 9, quotient=%d, remainder=%d\n", quotient, remainder)
	fmt.Println("Is 83 a prime no? :", math.IsPrime(83))

	hi := func() {
		fmt.Println("Hi")
	}
	hi()

	func() {
		fmt.Println("Hello")
	}()

	fmt.Println(utils.Counter())
	fmt.Println(utils.Counter())
	fmt.Println(utils.Counter())
}

func logOperation(operation func(int, int) int) func(int, int) int {
	return func(x int, y int) int {
		fmt.Printf("Initiating an opeation for %d and %d\n", x, y)
		result := operation(x, y)
		fmt.Printf("Returning result from the opeation %d\n", result)
		return result
	}
}
