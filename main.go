package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hi Chi!!"))
}

func main() {
	r := chi.NewRouter()
	r.Get("/", rootHandler)
	http.ListenAndServe(":3333", r)
}
