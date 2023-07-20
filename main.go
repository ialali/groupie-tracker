package main

import (
	"fmt"
	global "groupie-tracker/globals"
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	global.Client = &http.Client{}
	fs := http.FileServer(http.Dir("static"))
	handlers.GetJson("https://groupietrackers.herokuapp.com/api/artists", &global.Artists)
	handlers.GetJson("https://groupietrackers.herokuapp.com/api/locations", &global.Locations)
	handlers.GetJson("https://groupietrackers.herokuapp.com/api/dates", &global.Dates)
	handlers.GetJson("https://groupietrackers.herokuapp.com/api/relation", &global.Relations)
	global.AppendtoStruct()
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.Homepage)
	http.HandleFunc("/artists/", handlers.ArtistHandler)
	// http.HandleFunc("/artists/", idHandler) // funciona asi
	fmt.Println("Server started. Listening on http://localhost:8027")
	log.Fatal(http.ListenAndServe(":8027", nil))
}
