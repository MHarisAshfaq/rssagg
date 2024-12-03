package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// json.NewEncoder(w).Encode(data)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5XX error: ", message)
		// http.Error(w, message, http.StatusInternalServerError)
		// return
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errResponse{Error: message})
}
