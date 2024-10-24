package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	inputFile, inputErr := os.Open("input.txt")
	outputFile, outputErr := os.Create("output.txt")

	if inputErr != nil {
		fmt.Println(inputErr)
	}
	defer inputFile.Close()

	if outputErr != nil {
		fmt.Println(outputErr)
		os.Exit(1)
	}
	defer outputFile.Close()

	inputString := ""
	data := make([]byte, 1024)
	for {
		n, err := inputFile.Read(data)
		if err == io.EOF {
			break
		}
		inputString = string(data[:n])
	}
	inputArr := strings.Split(inputString, "\n")

	stringCounter := make(map[string]int)
	for i := 0; i < len(inputArr); i++ {
		stringCounter[inputArr[i]]++
	}

	individualStrings := []string{}
	for str, count := range stringCounter {
		if count == 1 {
			individualStrings = append(individualStrings, str)
		}
	}

	outputArr := make([]string, len(individualStrings))
	for i := 0; i < len(individualStrings); i++ {
		outputArr[i] = fmt.Sprintf(strings.ToUpper(individualStrings[i])+" - "+"%d "+"байт", len(individualStrings[i]))
	}
	sort.Strings(outputArr)

	for i := range outputArr {
		outputFile.WriteString(outputArr[i] + "\n")
	}
}
