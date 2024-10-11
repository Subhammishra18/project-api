package main

import (
	"context"
	"log"
	"net/http"

	"api/configuration" // Import the configuration package
	"api/controller"    // Import the controller package

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	configuration.Connect()
	defer configuration.Client.Disconnect(context.TODO()) // Ensure we disconnect when done

	// Initialize the router
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterUser).Methods("POST") // Handle registration
	r.HandleFunc("/login", controller.LoginUser).Methods("POST")       // Handle login

	// Start the server
	log.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r)) // Listen on port 8000
}
