package main

import (
	"net/http"
)

func main() {
 	// Lag en ny request multiplexer
 	// Motta kommende requests og send ut til matchende handlers
 	mux := http.NewServeMux()

	// Registrer endpoints og handlers
 	mux.Handle("/", &homeHandler{})

 	// Kjør serveren
 	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

// homeHandler sin ServeHTTP method
func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 	w.Write([]byte("Hello World!"))
}