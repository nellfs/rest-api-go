package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Picture struct {
	Msg string `json:"msg"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var picture Picture
	err := decoder.Decode(&picture)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
	}

	response := Picture{
		Msg: picture.Msg,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error Creating JSON response", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}
