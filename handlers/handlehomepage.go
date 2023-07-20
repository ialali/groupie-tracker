package handlers

import (
	"net/http"
	"text/template"
	"groupie-tracker/globals"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	data := map[string]interface{}{
		"Artists": global.Artists,
	}
	// Serve the HTML page with the filtered artists
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		return
	}
}
