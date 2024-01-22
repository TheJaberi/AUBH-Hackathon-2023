package AMJ

import (
	"html/template"
	"net/http"
)

// PageData represents the data used in the template for rendering the page.
type PageData struct {
	ShowOutput bool
	Output     string
	Str        string
	keyStr     string
}

// Handler handles the HTTP requests for the root path ("/").
func Handler(w http.ResponseWriter, r *http.Request) {
	// Check if the request path is not the root path
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}

	// Handle GET requests
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
