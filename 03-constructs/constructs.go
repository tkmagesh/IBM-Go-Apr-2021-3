package main

import (
	"fmt"
)

func main() {
	//if else
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	/* variable 'no' is not accessible outside the 'if' construct */
	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	//for
	//use as a 'for' construct
	/*
		sum := 0
		for i := 1; i <= 10; i++ {
			sum += i
		}
		fmt.Printf("Sum of the first 10 integers = %d\n", sum)
	*/

	//use as a 'while' construct
	/*
		sum := 1
		for sum < 100 {
			sum += sum
		}
		fmt.Println(sum)
	*/

	//use for an infinite loop
	sum := 1
	for {
		if sum > 100 {
			break
		}
		sum += sum
	}
	fmt.Println(sum)

	no := 21
	switch no % 2 {
	case 0:
		fmt.Printf("%d is an even number\n", no)
	case 1:
		fmt.Printf("%d is an odd number\n", no)
	}

	/*
		score
			0 - 3 => "Terrible"
			4 - 7 => "Not bad!"
			8 - 9 => "Good"
			10 => "Excellent"
			< 0 OR > 10 => "Unknown score"
	*/
	score := 11
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not bad!")
	case 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Unknown score")
	}

	var n int
	fmt.Println("Enter any value from 0 to 6:")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
	}
}
