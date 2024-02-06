package main

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/http"
)

// TODO
// Look into how to deal with this without them having to be outside of program
// flow
type Store = map[string]string

var store Store = make(map[string]string)

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	if url == "" {
		hash := r.URL.String()
		key := hash[1:]

		// TODO
		// Update store to work as a struct for moving over to postgres
		value, ok := store[key]

		if ok {
			w.Write([]byte(value))
		} else {
			// TODO
			// Move out http handlers
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 Not Found"))
		}
	} else {
		// TODO
		// Move hasher out so I can update how we hash in a single place
		// Maybe like a hash, unhash functions
		hasher := sha256.New()
		hasher.Write([]byte(url))
		bhash := hasher.Sum(nil)
		shash := base64.StdEncoding.EncodeToString(bhash)
		key := shash[:7]

		// TODO
		// Update store to work as a struct for moving over to postgres
		store[key] = url

		w.Write([]byte(key))
	}
}

func main() {
	mux := http.NewServeMux()

	// TODO
	// Move variable settings to a config file
	port := ":8080"

	mux.Handle("/", http.HandlerFunc(handler))

	// TODO
	// Create some logging functionality
	log.Print("Listening...")

	http.ListenAndServe(port, mux)
}
