package main

import (
	"fmt"
	"forum_reboot/server"
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	fmt.Println("loading server...")

	// Serve static HTML files
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("pages/"))))

	// Handle routes
	http.HandleFunc("/register", server.RegisterUser)
	http.HandleFunc("/login", server.LoginUser)
	http.HandleFunc("/logout", server.LogoutUser)
	http.HandleFunc("/is-logged-in", server.IsLoggedIn)
	// get posts
	http.HandleFunc("/get-posts", server.GetPosts)
	http.HandleFunc("/get-post", server.GetPost)
	http.HandleFunc("/create-post", server.CreatePost)
	http.HandleFunc("/add-comment", server.CreateComment)
	http.HandleFunc("/get-categories", server.GetCategories)
	http.HandleFunc("/categories", server.GetCategories) 

	// Handle like/dislike updates
	http.HandleFunc("/update-post-like", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	})
	


	// Redirect root to main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/pages/main.html", http.StatusSeeOther)
	})

	// Health check route
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		err := server.DB.Ping()
		if err != nil {
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Database is up and running"))
	})

	// Test database route
	http.HandleFunc("/testdb", func(w http.ResponseWriter, r *http.Request) {
		_, err := server.DB.Exec("CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, message TEXT)")
		if err != nil {
			http.Error(w, "Failed to create test table", http.StatusInternalServerError)
			return
		}

		_, err = server.DB.Exec("INSERT INTO test (message) VALUES (?)", "Hello, world!")
		if err != nil {
			http.Error(w, "Failed to insert test record", http.StatusInternalServerError)
			return
		}

		var message string
		err = server.DB.QueryRow("SELECT message FROM test WHERE id = 1").Scan(&message)
		if err != nil {
			http.Error(w, "Failed to query test record", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Test record message: " + message))
	})

	// Start the server
	fmt.Println("Server started on :2345")
	log.Fatal(http.ListenAndServe(":2345", nil))
}


