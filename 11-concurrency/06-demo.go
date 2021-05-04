package main

import "fmt"

func sum(nos []int, resultCh chan int) {
	result := 0
	for _, no := range nos {
		result += no
	}
	resultCh <- result
}

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	firstSet := data[:len(data)/2]
	secondSet := data[len(data)/2:]
	resultCh := make(chan int)
	go sum(firstSet, resultCh)
	go sum(secondSet, resultCh)
	total := <-resultCh + <-resultCh
	fmt.Println("Total = ", total)

}
