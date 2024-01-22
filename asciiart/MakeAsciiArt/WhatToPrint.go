package AMJ

// WhatToPrint generates the ASCII representation of a specific letter in a word.
// It takes the index 'i' indicating the style of the letter, the 'word' for which the ASCII representation is generated,
// the 'filename' of the font file, and the 'res' string to append the ASCII representation.
// Returns the updated 'res' string and an error if any.
func WhatToPrint(i int, word, filename, res string) (string, error) {
	for _, letter := range word {
		line, err := GetLine(1+(int(letter)-32)*9+i, filename)
		res += line
		if err != nil {
			return "", err
		}
	}
	return res, nil
}
