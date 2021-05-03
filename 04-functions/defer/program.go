package main

import "fmt"

func main() {
	f2()

}

func f2() {
	fmt.Println("f2 triggered")

	defer func() {
		fmt.Println("f2 completed")
	}()

	f1()
}
func f1() {
	fmt.Println("f1 triggered")

	defer func() {
		fmt.Println("f1 completed - defer 1")
	}()
	defer func() {
		fmt.Println("f1 completed - defer 2")
	}()
	defer func() {
		fmt.Println("f1 completed - defer 3")
	}()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("hard exit")
			fmt.Println(r)
		}
	}()
	panic("Something went wrong!")
}
