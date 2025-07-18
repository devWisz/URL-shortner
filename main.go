package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type URL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
}

var urlDb = make(map[string]URL)

// ===== Logic Functions Stay Same =====

func short_url(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	hash_data := hasher.Sum(nil)
	hash_string := hex.EncodeToString(hash_data)
	return hash_string[:7]
}

func CreateUrl(OriginalURL string) string {
	shortURL := short_url(OriginalURL)
	id := shortURL
	urlDb[shortURL] = URL{
		ID:          id,
		OriginalURL: OriginalURL,
		ShortURL:    shortURL,
		CreatedAt:   time.Now(),
	}
	return shortURL
}

func getreverseURL(id string) (string, error) {
	url, exists := urlDb[id]
	if !exists {
		return "", errors.New("URL not found")
	}
	return url.OriginalURL, nil
}

// ====== NEW HANDLERS BELOW ======

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func shorturlhandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || data.URL == "" {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	shorturl := CreateUrl(data.URL)
	response := struct {
		Shorturl string `json:"short_url"`
	}{Shorturl: shorturl}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLhandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getreverseURL(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func main() {
	fmt.Println("🚀 URL Shortener Server started at http://localhost:3000")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", shorturlhandler)
	http.HandleFunc("/redirect/", redirectURLhandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

