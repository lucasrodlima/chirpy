package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerReadChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.ReadAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retrieving chirps from database", err)
		return
	}

	respondWithJson(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerReadChirp(w http.ResponseWriter, r *http.Request) {
	pathID := r.PathValue("chirpID")

	userID := uuid.MustParse(pathID)

	chirp, err := cfg.db.ReadChirp(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	respondWithJson(w, http.StatusOK, chirp)
}
