package main

import (
	"fmt"
	"log"
	"net/http"
	"forum/helpers/handlers" 
)

func main() {
	println("loading server...")

	// Serve CSS files
	http.Handle("/pages/style/", http.StripPrefix("/pages/style/", http.FileServer(http.Dir("pages/style/"))))

	// Handle routes
	//http.HandleFunc("/", handlers.MakeHandler(handlers.MainHandler))
	//http.HandleFunc("/register", handlers.RegisterUser)
	//http.HandleFunc("/login", handlers.LoginUser)
	//http.HandleFunc("/logout", handlers.LogoutUser)

	// Start the server
	fmt.Println("Server started on :2345")
	log.Fatal(http.ListenAndServe(":2345", nil))
}
