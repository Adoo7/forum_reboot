package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userData struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Handle registration logic here
		fmt.Fprintf(w, "User registered: %s", userData.Username)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// LoginUser handles user login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Handle login logic here
		fmt.Fprintf(w, "User logged in: %s", credentials.Email)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// LogoutUser handles user logout
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var sessionData struct {
			SessionID string `json:"session_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&sessionData); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Handle logout logic here
		fmt.Fprintf(w, "User logged out, session ID: %s", sessionData.SessionID)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// UseHandler wraps a handler function
func UseHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}
