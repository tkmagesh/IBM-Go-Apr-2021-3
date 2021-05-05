package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileContents, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("Unable to read the file : %v", err)
	}
	fmt.Println(string(fileContents))
}
