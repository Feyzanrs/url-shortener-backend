package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	log.Println("Server is listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// HomeHandler - Ana sayfa handler'ı
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URL Shortener API is up and running!"))
}
