package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error : %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload any) {
	w.Header().Add("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshaling json error response: %s", err)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}
