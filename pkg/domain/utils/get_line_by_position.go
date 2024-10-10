package utils

import (
	"bufio"
	"fmt"
	"os"
)

// getLineByPositionInner returns the line that contains the specified character position (zero-based index) from the file_supportive
func getLineByPositionInner(filePath string, position int) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("could not open file_supportive: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("could not close file_supportive err:" + err.Error())
		}
	}(file)

	scanner := bufio.NewScanner(file)
	currentPos := 0
	row := 1
	for scanner.Scan() {
		line := scanner.Text()
		lineLength := len(line) + 1 // +1 for the newline character
		if currentPos+lineLength > position {
			return row, nil
		}
		currentPos += lineLength
		row++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file_supportive: %v", err)
	}

	return 0, fmt.Errorf("position %d out of range", position)
}

func GetLineByPosition(filePath string, position int) int {
	line, err := getLineByPositionInner(filePath, position)
	if err != nil {
		panic(err)
	}

	return line
}
