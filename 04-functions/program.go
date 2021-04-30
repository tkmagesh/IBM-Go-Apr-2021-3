package main

import (
	"fmt"
	"func-demos/calculator"
	gt "func-demos/greeter"
)

func main() {
	gt.Greet("Magesh", "Have a good day!")

	fmt.Printf("Adding 10 and 20 results in %d\n", calculator.Add(10, 20))
	quotient, remainder := calculator.Divide(80, 9)
	fmt.Printf("Dividing 80 by 9, quotient=%d, remainder=%d\n", quotient, remainder)
}
