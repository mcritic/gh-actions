package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Vulnerability: HTTP response headers are being controlled by the user
		w.Header().Set("X-Custom-Header", r.URL.Query().Get("header_value"))
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", nil)
}
