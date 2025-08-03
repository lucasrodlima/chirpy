package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		fmt.Printf("Error decoding json: %v", err)

		type returnVals struct {
			Error string `json:"error"`
		}

		respBody := returnVals{
			Error: "Something went wrong",
		}

		data, err := json.Marshal(respBody)
		if err != nil {
			fmt.Printf("Error marshaling json error response: %s", err)
			return
		}

		w.WriteHeader(500)
		w.Write(data)
		return
	}

	if len(params.Body) > 140 {
		type returnVals struct {
			Error string `json:"error"`
		}

		respBody := returnVals{
			Error: "Chirp is too long",
		}

		data, err := json.Marshal(respBody)
		if err != nil {
			fmt.Printf("Error marshaling json error response: %s", err)
			return
		}

		w.WriteHeader(400)
		w.Write(data)
		return

	}

	type returnVals struct {
		Valid bool `json:"valid"`
	}

	respBody := returnVals{
		Valid: true,
	}

	data, err := json.Marshal(respBody)
	if err != nil {
		fmt.Printf("Error marshaling json error response: %s", err)
		return
	}

	w.WriteHeader(200)
	w.Write(data)
	return
}
