package models

// User represents the user data model
type User struct {
	Name        string `json:"name" bson:"name"`                 // Name of the user
	Email       string `json:"email" bson:"email"`               // Email of the user
	Age         int    `json:"age" bson:"age"`                   // Age of the user
	PhoneNumber string `json:"phone_number" bson:"phone_number"` // Phone number of the user
}
