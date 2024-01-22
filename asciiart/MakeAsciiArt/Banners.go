package AMJ

import (
	"errors"
	"strings"
)

// Banners retrieves the filename for the specified banner style.
// Returns the filename and an error if the style is not found.
func Banners(banner string) (string, error) {
	banners := map[string]string{
		"shadow":     "shadow.txt",
		"standard":   "standard.txt",
		"thinkertoy": "thinkertoy.txt",
		"doom":       "doom.txt",
		"extrafont":  "extrafont.txt",
		// Add more banners here as needed
	}

	// Convert the banner style to lowercase
	banner = strings.ToLower(banner)

	// Check if the banner style exists in the map
	if banner1, ok := banners[banner]; ok {
		style := "asciiart/Fonts/" + banner1
		return style, nil
	} else {
		return "", errors.New("500: Internal Server Error")
	}
}
