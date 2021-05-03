package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Sit adipisicing deserunt eiusmod mollit ut exercitation. Cupidatat exercitation et incididunt ex dolore irure deserunt reprehenderit id pariatur aliquip velit sit enim. Incididunt sit minim eiusmod excepteur amet commodo cillum fugiat. Consectetur et occaecat cupidatat enim amet laborum. Velit elit dolor velit quis qui non velit cillum pariatur nostrud dolor esse. Voluptate Lorem occaecat cillum ullamco proident magna deserunt. Commodo tempor aliquip veniam cupidatat ex nostrud ipsum nulla Lorem ipsum."
	data := strings.ReplaceAll(str, ".", "")
	words := strings.Split(data, " ")
	wordOccurancesBySize := map[int][]string{}
	for _, word := range words {
		wordSize := len(word)
		wordOccurancesBySize[wordSize] = append(wordOccurancesBySize[wordSize], word)
	}
	sizeOfMaxOccurance, maxOccurance := 0, 0
	for size, words := range wordOccurancesBySize {
		if len(words) > maxOccurance {
			maxOccurance = len(words)
			sizeOfMaxOccurance = size
		}
	}
	fmt.Printf("Size of the max occured words = %d, count = %d\n", sizeOfMaxOccurance, maxOccurance)
}

//group the words by their size
//find out the most occuring size of the string and print the count
