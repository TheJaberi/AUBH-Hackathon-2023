package AMJ

import (
	"strings"
)

// AlphabetFormat replaces all occurrences of the escape sequence "\t" in the given string with four spaces ("    ").
// Returns the modified string.
func AlphabetFormat(s string) string {
	s = strings.ReplaceAll(s, "\\t", "    ")
	return s
}
