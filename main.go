package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Artist struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    Locations
	Dates        Dates
	Relations    Relations
}
type Relation struct {
	Id             int                    `json:"id"`
	DatesLocations map[string]interface{} `json:"datesLocations"`
}
type Relations struct {
	Index []Relation `json:"index"`
}
type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Dates struct {
	Index []Date `json:"index"`
}
type Location struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
	Dates    string   `json:"dates"`
}
type Locations struct {
	Index []Location `json:"index"`
}

var (
	client    *http.Client
	locations Locations
	myartist  []Artist
	dates     Dates
	relations Relations
)

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	fs := http.FileServer(http.Dir("static"))
	GetJson("https://groupietrackers.herokuapp.com/api/artists", &myartist)
	GetJson("https://groupietrackers.herokuapp.com/api/locations", &locations)
	GetJson("https://groupietrackers.herokuapp.com/api/dates", &dates)
	GetJson("https://groupietrackers.herokuapp.com/api/relation", &relations)
	AppendtoStruct()
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", homepage)
	http.HandleFunc("/artists/", artistHandler)
	// http.HandleFunc("/artists/", idHandler) // funciona asi
	fmt.Println("Server started. Listening on http://localhost:8027")
	log.Fatal(http.ListenAndServe(":8027", nil))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	data := map[string]interface{}{
		"Artists": myartist,
	}
	// Serve the HTML page with the filtered artists
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		return
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
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
	var artist Artist
	for _, a := range myartist {
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
	err = tmpl.Execute(w, artist)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func AppendtoStruct() {
	for index := range locations.Index {
		myartist[index].Locations.Index = append(myartist[index].Locations.Index, locations.Index[index])
	}
	for index := range dates.Index {
		myartist[index].Dates.Index = append(myartist[index].Dates.Index, dates.Index[index])
	}
	for index := range relations.Index {
		myartist[index].Relations.Index = append(myartist[index].Relations.Index, relations.Index[index])
	}
}

func GetJson(url string, target interface{}) error {
	respose, err := client.Get(url)
	if err != nil {
		return err
	}
	defer respose.Body.Close()
	return json.NewDecoder(respose.Body).Decode(target)
}

// func idHandler(w http.ResponseWriter, r *http.Request) {
// 	data := myartist
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	path := r.URL.Path
// 	segments := strings.Split(path, "/artists/")
// 	userID := segments[len(segments)-1]
// 	id, err := strconv.Atoi(userID)
// 	if err != nil {
// 		// fmt.Fprintln(w, "Invalid ID")
// 		w.WriteHeader(400)
// 		w.Write([]byte("ID cold not be converted to Interger"))
// 		return
// 	}
// 	for _, task := range data {
// 		if task.Id == id {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(task)

// 		}
// 	}
// }
