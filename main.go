package main

import (
	"forum/helpers/handlers"
	"log"
	"net/http"
)

func main() {
	println("loading server...")

	//serve css file
	http.Handle("/pages/style/", http.StripPrefix("/pages/style/", http.FileServer(http.Dir("pages/style/"))))

	//handle routes
	http.HandleFunc("/", handlers.MakeHandler(handlers.MainHandler))
	// example of how to make a function return a response
	// http.HandleFunc("/get-artists", handlers.GetArtists)
	// http.HandleFunc("/get-relations", handlers.GetRelations)
	log.Fatal(http.ListenAndServe(":2345", nil))
}
