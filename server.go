package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type KeyData struct {
	Keys []string `json:"keys"`
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "piano_interval_fetch.html")
}
func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Decode JSON from request body
	var data KeyData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Received key data: %+v\n", data.Keys)
	cal_interval := Interval_main(data.Keys)
	fmt.Printf("Interval calculated: %+v\n", cal_interval)
	response := map[string]string{
		"message":  "Data received successfully!",
		"interval": cal_interval,
	}
	json.NewEncoder(w).Encode(response)
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/api/keydata", handlePost)
	fmt.Println("Server starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
