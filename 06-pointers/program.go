package main

import "fmt"

func main() {
	no := 10
	fmt.Println(no)
	var noPtr *int = &no
	fmt.Println(noPtr)

	//dereferencing the value from the address
	fmt.Println("Deferencing a pointer")
	fmt.Println(*noPtr)
	/* no == *(&no) */

	fmt.Println("Incrmenting operation")
	fmt.Printf("Before incrementing no = %d\n", no)
	incrment(&no)
	fmt.Printf("After incrementing no = %d\n", no)

	x, y := 10, 20
	fmt.Printf("Before swapping, x = %d and y = %d\n", x, y)
	swap( /*  */ )
	fmt.Printf("After swapping, x = %d and y = %d\n", x, y)
}

func incrment(x *int) {
	*x += 1
}

func swap( /*  */ ) {

}
