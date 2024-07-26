package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "GO")
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Displays a particular snippet with ID %d", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displays a form to create a snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("server is starting on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
