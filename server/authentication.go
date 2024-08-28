package server

import (

	"database/sql"



	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("sqlite3", "./forum.sqlite")
	if err != nil {
		panic(err)
	}
}
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("RegisterUser called")
	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Retrieve form values
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Check if email is already taken
		var exists bool
		err = DB.QueryRow("SELECT EXISTS(SELECT 1 FROM User WHERE email = ?)", email).Scan(&exists)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w, "Email already taken", http.StatusConflict)
			return
		}

		// Insert user into the database
		_, err = DB.Exec("INSERT INTO User (username, email, passwords) VALUES (?, ?, ?)", username, email, password)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}



func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Retrieve form values
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Log the input email and password for debugging
		log.Printf("Email: %s, Password: %s", email, password)

		// Retrieve stored password from the database
		var storedPassword string
		err = DB.QueryRow("SELECT passwords FROM User WHERE email = ?", email).Scan(&storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Log the stored password for debugging
		log.Printf("Stored Password: %s", storedPassword)

		// Check if the provided password matches the stored password
		if password != storedPassword {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Create session
		sessionID := uuid.New().String()
		expireTime := time.Now().Add(1 * time.Hour) // Session expires in 1 hour
		_, err = DB.Exec("INSERT INTO UserSession (UserSessionID, User_ID, Token, ExpireTime) SELECT ?, User_ID, ?, ? FROM User WHERE email = ?", sessionID, sessionID, expireTime, email)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Set cookie
		http.SetCookie(w, &http.Cookie{
			Name:    "session_id",
			Value:   sessionID,
			Expires: expireTime,
			Path:    "/",
		})

		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}



func IsLoggedIn(w http.ResponseWriter, r *http.Request) {
	_, loggedIn := CheckSession(r)
	if loggedIn {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}



func CheckSession(r *http.Request) (int, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return 0, false
	}

	var userID int
	err = DB.QueryRow("SELECT User_ID FROM UserSession WHERE Token = ? AND ExpireTime > ?", cookie.Value, time.Now()).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, false
		}
		return 0, false
	}

	return userID, true
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "No active session", http.StatusUnauthorized)
		return
	}

	// Delete session from the database
	_, err = DB.Exec("DELETE FROM UserSession WHERE Token = ?", cookie.Value)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Expire the cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
		Path:    "/",
	})

	w.WriteHeader(http.StatusOK)
}

// server/authentication.go
func init() {
	var err error
	DB, err = sql.Open("sqlite3", "./forum.sqlite")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")
}
