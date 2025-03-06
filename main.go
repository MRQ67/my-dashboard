package main

import (
	"dashboard/tools"
	"encoding/json"
	"log"
	"net/http"
)

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City is required", http.StatusBadRequest)
		return
	}

	/*API_KEY := os.Getenv("WEATHER_API_KEY")
	if API_KEY == "" {
		fmt.Println("Error: WEATHER_API_KEY environment variable not set")
		os.Exit(1)
	}*/

	data, err := tools.GetWeather(city)
	if err != nil {
		http.Error(w, "Failed to get weather", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortURL, err := tools.ShortenURL(longURL)
	if err != nil {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}

func githubHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	profile, err := tools.GetGitHubProfile(username)
	if err != nil {
		http.Error(w, "Failed to get GitHub profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Add weather API endpoint
	http.HandleFunc("/weather", weatherHandler)
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/github", githubHandler)

	// Start the server
	log.Println("Server starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
