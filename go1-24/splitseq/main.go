package main

import (
	//"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {

	csvFile := "data.csv"
	SEPARATOR := ","

	csvBytes, err := os.ReadFile(csvFile)
	if err != nil {
		panic(err)
	}

	csvString := string(csvBytes)
	for _, csvLine := range strings.SplitAfter(csvString, "\n") {
		for _, csvField := range strings.Split(csvLine, SEPARATOR) {
			fmt.Println("Field", csvField)
		}
	}
}
