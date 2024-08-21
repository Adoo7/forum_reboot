package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3" /
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		panic(err)
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if email is already taken
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM User WHERE email = ?)", user.Email).Scan(&exists)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Email already taken", http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Insert user into the database
	_, err = db.Exec("INSERT INTO User (username, email, passwords) VALUES (?, ?, ?)", user.Username, user.Email, string(hashedPassword))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Retrieve hashed password from the database
	var hashedPassword string
	err := db.QueryRow("SELECT passwords FROM User WHERE email = ?", credentials.Email).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Compare provided password with stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Create session
	sessionID := uuid.New().String()
	expireTime := time.Now().Add(1 * time.Hour) // Session expires in 1 hour
	_, err = db.Exec("INSERT INTO UserSession (UserSessionID, User_ID, Token, ExpireTime) SELECT ?, User_ID, ?, ? FROM User WHERE email = ?", sessionID, sessionID, expireTime, credentials.Email)
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
}

func CheckSession(r *http.Request) (int, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return 0, false
	}

	var userID int
	err = db.QueryRow("SELECT User_ID FROM UserSession WHERE Token = ? AND ExpireTime > ?", cookie.Value, time.Now()).Scan(&userID)
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
	_, err = db.Exec("DELETE FROM UserSession WHERE Token = ?", cookie.Value)
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

