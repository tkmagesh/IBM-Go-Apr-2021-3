package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNos := getRandomNos()
	//primes := filterPrimes(randomNos)
	primes := filterNos(randomNos, isPrime)
	fmt.Println("Primes = ", primes)
	fmt.Println("Prime Count = ", len(primes))

	evenNos := filterNos(randomNos, isEven)
	fmt.Println("Even Nos = ", evenNos)
	fmt.Println("Even Nos Count = ", len(evenNos))
}

func getRandomNos() []int {
	nos := [50]int{}
	randomizer := rand.New(rand.NewSource(2))
	for idx, _ := range nos {
		nos[idx] = randomizer.Intn(200)
	}
	return nos[:]
}

func isPrime(n int) bool {
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/*
func filterPrimes(nos []int) []int {
	primes := []int{}
	for _, no := range nos {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}
*/

func isEven(n int) bool {
	return n%2 == 0
}

/*
func filterEven(nos []int) []int {
	evenNos := []int{}
	for _, no := range nos {
		if isEven(no) {
			evenNos = append(evenNos, no)
		}
	}
	return evenNos
}
*/

func filterNos(nos []int, criteriaFn func(int) bool) []int {
	resultList := []int{}
	for _, no := range nos {
		if criteriaFn(no) {
			resultList = append(resultList, no)
		}
	}
	return resultList
}
