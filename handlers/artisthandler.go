package handlers

import (
	global "groupie-tracker/globals"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del artista desde la URL
	path := r.URL.Path
	segments := strings.Split(path, "/artists/")
	if len(segments) != 2 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	userID := segments[len(segments)-1]
	// artistIDStr := segments[2]
	artistID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid Artist ID", http.StatusBadRequest)
		return
	}

	// Buscar el artista por su ID
	var artist global.Artist // Declare the artist variable here
	for _, a := range global.Artists {
		if a.Id == artistID {
			artist = a
			break
		}
	}

	// Verificar si se encontró el artista
	if artist.Name == "" {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	// Cargar la plantilla HTML
	tmpl, err := template.ParseFiles("template/artist.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Renderizar la plantilla con los datos del artista
	err = tmpl.Execute(w, artist) // Use the 'artist' variable here
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
