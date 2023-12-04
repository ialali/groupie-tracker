package handlers

import (
	global "groupie-tracker/globals"
	"net/http"
	"text/template"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	// Check if the request is a form submission
	if r.Method == http.MethodGet {
		// Get the search query from the form submission
		query := r.URL.Query().Get("query")

		// Use the PerformSearch function to perform the search
		results := PerformSearch(query)

		// Create a map with the search results to pass to the template
		data := map[string]interface{}{
			"Artists": results,
		}

		// Serve the HTML page with the search results
		tmpl := template.Must(template.ParseFiles("index.html"))
		err := tmpl.Execute(w, data)
		if err != nil {
			return
		}
		return
	}

	// If it's not a form submission, serve the default page with all artists
	data := map[string]interface{}{
		"Artists": global.Artists,
	}

	// Serve the HTML page with the default artists
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
