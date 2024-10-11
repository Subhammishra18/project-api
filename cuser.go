package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"api/configuration" // Import the configuration package
	"api/models"        // Import the models package

	"go.mongodb.org/mongo-driver/bson"
)

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User // Create a User object

	// Decode the incoming JSON request into the User object
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the user into the MongoDB collection
	collection := configuration.Client.Database("testdb").Collection("users")
	if _, err := collection.InsertOne(context.TODO(), user); err != nil {
		http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response back to the client
	w.WriteHeader(http.StatusCreated) // 201 Created
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

// LoginUser handles user login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User // Create a User object

	// Decode the incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Look for the user by email in the MongoDB collection
	collection := configuration.Client.Database("testdb").Collection("users")
	var foundUser models.User
	if err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&foundUser); err != nil {
		http.Error(w, "User not found: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Send a response back to the client
	w.WriteHeader(http.StatusOK) // 200 OK
	if err := json.NewEncoder(w).Encode(foundUser); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
