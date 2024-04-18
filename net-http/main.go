package main

import (
	"net/http"
)

// PARAMS

func getItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("received request for item: " + id))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/item/{id}", getItem)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
