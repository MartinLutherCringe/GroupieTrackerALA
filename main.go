package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Fetch API data by URL
	fetchData("https://groupietrackers.herokuapp.com/api/artists", &artistsList)
	fetchData("https://groupietrackers.herokuapp.com/api/relation", relationList)

	// URL manager for "/"
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/artists", artistIndexHandler)

	// File server for CSS files
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	// get port number from "PORT" env variable
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080" // Use default "8080" port number if env variable is not defined
	}

	// Start server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) { // URL manager for "/"
	logRequest(r) // Enregistrer la requête
	if r.URL.Path != "/" || r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound) // Manage error request
		return
	}

	writeTemplate(w, "index.html", artistsList)
}

func artistIndexHandler(w http.ResponseWriter, r *http.Request) { // URL manager for "/artists"
	if r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound) // Manage request errors
		return
	}

	id, err := extractQueryID(w, r) // Extract ID request
	if err != nil {
		log.Println(err.Error())
		errorHandler(w, r, http.StatusNotFound) // Manage request errors
		return
	}

	data, err := getArtistByID(id) // Fetch artist data corresponding to ID
	if err != nil {
		log.Println(err.Error())
		errorHandler(w, r, http.StatusNotFound) // Manage request errors
		return
	}

	writeTemplate(w, "artist.html", data)
}