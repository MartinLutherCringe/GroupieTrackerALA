package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

var artistsList = []artist{}
var relationList = &relation{}

func fetchData(api string, data interface{}) { // Function who fetch data to API
	log.Println("Début de la synchronisation de l'API " + api)
	res, err := http.Get(api) // GET Request to API
	if err != nil {
		log.Println(err.Error())
		return
	}

	if res.StatusCode == http.StatusOK { // If request is successful
		bodyBytes, err := io.ReadAll(res.Body) // Read body response
		res.Body.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}

		err = json.Unmarshal(bodyBytes, &data) // Organization of JSON data in structure
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	log.Println("Synchronisation de l'API effectuée " + api)
}

func getArtistByID(id int) (*artistData, error) { // Find artist by ID
	artist := filterArtistByID(artistsList, id)
	if artist == nil {
		return nil, errors.New("Pas d'artistes trouvé..") // Return error if no artist is found
	}

	var data = artistData{Artist: *artist}
	var dates = filterRelationByID(relationList.Index, id)
	if dates != nil {
		data.DatesLocations = make(map[string]interface{})
		for key, value := range dates.DatesLocations {
			var locationName = strings.ReplaceAll(key, "_", " ") // Format the name of location
			locationName = strings.ReplaceAll(locationName, "-", " - ")
			locationName = strings.ToUpper(locationName)
			data.DatesLocations[locationName] = value
		}
	}

	return &data, nil
}

func filterArtistByID(artists []artist, id int) *artist {
	for _, item := range artists {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func filterRelationByID(relations []artistRelation, id int) *artistRelation {
	for _, item := range relations {
		if item.ID == id {
			return &item
		}
	}
	return nil
}
