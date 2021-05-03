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
	swap(&x, &y)
	fmt.Printf("After swapping, x = %d and y = %d\n", x, y)

	nos := []int{10, 20, 30}
	fmt.Println("Before appending a value to the slice ", nos)
	addValue(&nos, 40)
	fmt.Println("After appending 40 to the slice ", nos)

	fmt.Println("Before updating a value in the slice ", nos)
	updateValue(nos, 0, 100)
	fmt.Println("After updating a value in the slice ", nos)
}

func updateValue(list []int, idx int, value int) {
	list[idx] = value
}

func addValue(list *[]int, value int) {
	*list = append(*list, value)
}

func incrment(x *int) {
	*x += 1
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
