package AMJ

// AllEmpty checks if all strings in the given array are empty.
// Returns true if all strings are empty, false otherwise.
func AllEmpty(arr []string) bool {
	for _, word := range arr {
		if word != "" {
			return false
		}
	}
	return true
}
