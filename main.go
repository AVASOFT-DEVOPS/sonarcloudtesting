package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

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

		validate := validate(name)

		if validate != nil {
			fmt.Fprintf(w, validate.Message)
			return
		}

		// Write the response back to the client
		fmt.Fprintf(w, "Hi, %s! How you doing?", name)
	}).Methods("GET")

	// Start the HTTP server
	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

type ValidateResponse struct {
	Code    string
	Message string
}

func validate(req string) *ValidateResponse {

	alphabetsRegex := regexp.MustCompile(`^[a-zA-Z]*$`)
	if strings.TrimSpace(req) == "" {
		return &ValidateResponse{Code: "400", Message: "Name cannot be empty"}
	} else if !alphabetsRegex.MatchString(req) {
		return &ValidateResponse{Code: "400", Message: "Name should not contain numbers"}
	} else {
		return nil
	}

}
