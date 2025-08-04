package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/lucasrodlima/chirpy/internal/database"
)

func (cfg *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserId string `json:"user_id"`
		Body   string `json:"body"`
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			"Error decoding parameters", err)
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", err)
		return
	}

	cleanChirp := cleanMessage(params.Body)

	userId := uuid.MustParse(params.UserId)

	newChirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   cleanChirp,
		UserID: userId,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			"Couldn't create new chirp, maybe non-existent user id", err)
		return
	}

	respondWithJson(w, http.StatusCreated, newChirp)
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
