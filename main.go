package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
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

func short_url(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("Hasher:", hasher)
	hash_data := hasher.Sum(nil)
	fmt.Println("Hash Data:", hash_data)
	hash_string := hex.EncodeToString(hash_data)
	fmt.Println("Hash:", hash_string[:7])

	return hash_string[:7]
}

func CreateUrl(OriginalURL string) string {
	shortURL := short_url(OriginalURL)
	id := shortURL
	urlDb[shortURL] = URL{
		ID:          id,
		OriginalURL: OriginalURL,
		ShortURL:    shortURL,
		CreatedAt:   time.Now()}
	return shortURL
}

func getreverseURL(id string) (string, error) {
	url, exists := urlDb[id]
	if !exists {
		return url.ID, errors.New("URL not found in this ID")
	}
	return url.OriginalURL, nil
}

// handle func
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Writing my first web based URL ")
}

func redirectURLhandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getreverseURL(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func shorturlhandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	shorturl := CreateUrl(data.URL)
	//fmt.Fprintf(w, "Short URL is: %s",shorturl)
	response := struct {
		Shorturl string `json:"short_url"`
	}{
		Shorturl: shorturl,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	fmt.Println("Starting the url shortner project")
	OriginalURL := "https://www.youtube.com/@LearnEassy"
	CreateUrl(OriginalURL)

	//handler function part
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", shorturlhandler)

	http.HandleFunc("/redirect/", redirectURLhandler)

	//starting the server on localhost:3000
	fmt.Println("Starting the server on localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error to starting  the server", err)

	}
}
