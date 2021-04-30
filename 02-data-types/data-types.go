package main

import "fmt"

func main() {
	//x, y := 100, 200
	const x, y uint8 = 100, 200
	//x = x + 100
	sum := uint16(x) + uint16(y)
	fmt.Printf("Sum of %d (%T) and %d (%T) is %d (%T)\n", x, x, y, y, sum, sum)
}
