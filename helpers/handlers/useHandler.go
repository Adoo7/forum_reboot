package handlers

import (
	"fmt"
	"net/http"
)

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle registration logic here
		fmt.Fprintln(w, "User registered successfully")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// LoginUser handles user login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle login logic here
		fmt.Fprintln(w, "User logged in successfully")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// LogoutUser handles user logout
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle logout logic here
		fmt.Fprintln(w, "User logged out successfully")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// MakeHandler wraps a handler function
func MakeHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}
