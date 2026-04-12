package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from handler!"))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
