package AMJ

import (
	"errors"
	"strings"
)

// Asciiart generates ASCII art based on the provided string and font.
// Returns the generated ASCII art as a string and an error if any.
func Asciiart(str, font string) (string, error) {
	// Retrieve the filename for the specified font
	filename, err := Banners(font)
	if err != nil {
		return "", err
	}

	restoprint := "" // Stores the output file

	// Replace "\r\n" with "\\n" to ensure compatibility with the ASCII art program
	strarr := strings.Split(str, "\r\n")
	str = strings.Join(strarr, "\\n")

	// Check if the input string contains non-English characters
	for _, r := range str {
		if r < 32 || r > 126 {
			return "", errors.New("400: Bad Request: Use English chars only")
		}
	}

	str = AlphabetFormat(str) // Fix the "\t" characters

	res := ""
	args := strings.Split(str, "\\n")

	// Check if the input consists of empty lines only
	if AllEmpty(args) {
		args = args[1:]
	}

	// Write text line by line into 'res' and store it in 'restoprint'
	for _, word := range args {
		if word == "" {
			restoprint = restoprint + "\n"
			continue
		}

		for i := 0; i < 8; i++ {
			res, err = WhatToPrint(i, word, filename, res)
			if err != nil {
				return "", err
			}
			restoprint = restoprint + "\n" + res
			res = ""
		}
	}

	return restoprint[1:], nil
}
