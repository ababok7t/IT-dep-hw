package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var filePATH string

func main() {
	fmt.Fscan(os.Stdin, &filePATH)
	inputFile, inputErr := os.Open(string(filePATH))
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
	if (len(outputArr) == 1) && (outputArr[0] == " - 0 байт") {
		outputFile.WriteString("файл пуст")
		fmt.Println("файл пуст")
	} else {
		for i := range outputArr {
			outputFile.WriteString(outputArr[i] + "\n")
		}
	}
}
