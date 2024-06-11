package main

// Structure for artists (based on API structure)
type artist struct {
	ID           int
	Name         string
	CreationDate int
	FirstAlbum   string
	Image        string
	Members      []string
}

// Structure for relation between artists
type relation struct {
	Index []artistRelation
}

// Structure for relation of artist, dates and locations
type artistRelation struct {
	ID             int
	DatesLocations map[string]interface{} // Map dates and associated location (key: name of loaction, value: data)
}

// Structure for artist data, y compris les dates et les lieux associés
type artistData struct {
	Artist         artist                  // Artiste
	DatesLocations map[string]interface{} // Map des dates et des lieux associés (clé: nom de la localisation, valeur: données supplémentaires)
}