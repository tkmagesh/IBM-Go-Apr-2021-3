package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileHandle, fileError := os.Open("emp.csv")
	if fileError != nil {
		log.Fatalf("Error opening the file %v\n", fileError)
	}
	defer fileHandle.Close()
	reader := csv.NewReader(fileHandle)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading the string %v :", err)
		}
		fmt.Println(line[0], line[1], line[2])
	}
}
