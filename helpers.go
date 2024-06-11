package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) { // Error manager
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		writeTemplate(w, "NotFound.html", "404 Not Found")
	}
}

func logRequest(r *http.Request) { // FUnction for log requests in console
	log.Printf("%v %v requested", r.Method, r.URL.Path)
}

func writeTemplate(w http.ResponseWriter, templateName string, data interface{}) { // Print content to webpage
	t, err := template.ParseFiles(templateName)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = t.Execute(w, data) // Execute template with provided data
	if err != nil {
		log.Fatal(err.Error())
	}
}

func extractQueryID(w http.ResponseWriter, r *http.Request) (int, error) { // Extract ID of request
	keys, ok := r.URL.Query()["ID"]
	if !ok || len(keys) != 1 {
		return 0, errors.New("L'URL ID est manquant") // Return error if "ID" key is missing or if there are severals values
	}
	key := keys[0]
	id, err := strconv.Atoi(key) // Convert value of ID to int value
	if err != nil {
		return 0, err
	}

	return id, nil
}