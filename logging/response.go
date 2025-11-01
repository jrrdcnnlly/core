package logging

import "net/http"

// Wraps http.ResponseWriter to record response status code.
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Create a new ResponseWriter.
func newResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w, http.StatusOK}
}

// Write the response status code.
func (w *ResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
