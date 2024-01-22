package AMJ

import (
	"net/http"
)

// CSSHandler handles HTTP requests for serving CSS files.
func CSSHandler(w http.ResponseWriter, r *http.Request) {
	// Construct the file path by appending "views" to the URL path
	filePath := "views" + r.URL.Path

	// Serve the file using http.ServeFile
	http.ServeFile(w, r, filePath)
}
