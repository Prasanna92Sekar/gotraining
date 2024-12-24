package handlers

import (
	"encoding/json"
	"gotraining/models"
	"net/http"

	"github.com/go-chi/chi"
)

// the users whatever we created, that will be stored in 'users'
// dict of dicts
// need to configure database, and need to add details over there
var users = make(map[string]models.User)

// FUnction to create a User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Checking whether the given input is valid or not
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	// Handling mandatory fields by checking its data
	// Here Mobile num is not a mandatory field
	if user.ID == "" || user.Name == "" || user.Email == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}
	// Creating the User
	users[user.ID] = user
	w.WriteHeader(http.StatusCreated)                  // HTTP Status code (Success-201)
	w.Header().Set("Content-Type", "application/json") // Content type (JSON)
	json.NewEncoder(w).Encode(user)                    // Encoding the JSON response
}

// Function to get all the Users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	usersList := []models.User{} // Initializing the users list for storing the user data
	// looping the users
	for _, user := range users {
		usersList = append(usersList, user) // Appending each user to the users list
	}
	w.Header().Set("Content-Type", "application/json") // Content type (JSON)
	json.NewEncoder(w).Encode(usersList)               // Encoding the JSON response
}

// Function to get User data based on ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // Getting the user id from API end point (URL Params)
	user, exists := users[id]   // Getting user data
	// Check whether the user is existing or not
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		// User not found msg with HTTP Status code (NotFound-404)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
