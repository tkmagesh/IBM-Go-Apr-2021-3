package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	data := `
firstname, lastname, age
magesh, kuppan, 42
suresh, kannan, 40
ramesh, jayaraman, 43
`
	reader := csv.NewReader(strings.NewReader(data))
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
