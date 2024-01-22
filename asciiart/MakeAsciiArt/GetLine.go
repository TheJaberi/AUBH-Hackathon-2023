package AMJ

import (
	"bufio"
	"errors"
	"os"
)

// GetLine retrieves a specific line from a file identified by the filename.
// Returns the requested line and an error if any.
func GetLine(num int, filename string) (string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil { // Check if there was an error opening the file
		return "", errors.New("500: Internal Server Error: Error opening file: " + err.Error())
	}
	defer file.Close()

	str := ""
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	// Scan through the file until the desired line number is reached
	for scanner.Scan() {
		if lineNumber == num {
			str = scanner.Text()
		}
		lineNumber++
	}

	return str, nil
}
