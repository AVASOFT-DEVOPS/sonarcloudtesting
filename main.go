package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router instance
	r := mux.NewRouter()

	// Define the "/hello/{name}" endpoint
	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		// Extract the "name" path parameter from the request
		vars := mux.Vars(r)
		name := vars["name"]

		// Write the response back to the client
		fmt.Fprintf(w, "Hi, %s! How you doing?", name)
	}).Methods("GET")

	// Start the HTTP server
	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
