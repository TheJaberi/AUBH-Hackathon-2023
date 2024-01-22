package AMJ

import (
	"net/http"
)

// BadRequestErrorHandler handles HTTP requests for a Bad Request (400) error.
// It writes the appropriate status code, and serves the corresponding error page based on whether the "textbox" form value is empty or not.
func BadRequestErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	if r.FormValue("textbox") != "" {
		http.ServeFile(w, r, "views/errorpages/eng400.html")
	} else {
		http.ServeFile(w, r, "views/errorpages/400.html")
	}
}

// NotFoundHandler handles HTTP requests for a Not Found (404) error.
// It writes the appropriate status code and serves the corresponding error page.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "views/errorpages/404.html")
}

// InternalServerErrorHandler handles HTTP requests for an Internal Server Error (500) error.
// It writes the appropriate status code and serves the corresponding error page.
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	http.ServeFile(w, r, "views/errorpages/500.html")
}
