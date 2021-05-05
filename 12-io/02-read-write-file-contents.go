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
	if err := ioutil.WriteFile("sample-copy.txt", fileContents, 0x777); err != nil {
		log.Fatalf("Unable to write the file : %v", err)
	} else {
		fmt.Println("File contents copied")
	}
}
