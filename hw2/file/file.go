package file

import (
	"os"
)

func ReadFile(path string) (string, error) {
	inputFile, inputErr := os.Open(path)
	if inputErr != nil {
		return "", inputErr
	}
	defer inputFile.Close()
	data := make([]byte, 1024)
	readingByte, readingErr := inputFile.Read(data)
	if readingErr != nil {
		return "", readingErr
	}
	text := string(data[:readingByte])
	return text, nil
}
