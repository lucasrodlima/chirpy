package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			"couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", err)
		return
	}

	cleanChirp := cleanMessage(params.Body)

	respondWithJson(w, http.StatusOK, returnVals{
		CleanedBody: cleanChirp,
	})
}

func cleanMessage(msg string) string {
	msgWords := strings.Split(msg, " ")
	badWords := []string{"kerfuffle", "sharbert", "fornax"}

	for i, word := range msgWords {
		for _, badWord := range badWords {
			if strings.ToLower(word) == badWord {
				msgWords[i] = "****"
			}
		}
	}
	cleanStr := strings.Join(msgWords, " ")
	return cleanStr
}
