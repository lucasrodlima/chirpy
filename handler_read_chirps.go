package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerReadChirps(w http.ResponseWriter, r *http.Request) {
	if q := r.URL.Query().Get("author_id"); q != "" {
		userID, err := uuid.Parse(q)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "could parse author id", err)
			return

		}

		chirps, err := cfg.db.ReadChirpsFromUser(r.Context(), userID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error retrieving chirps from database", err)
			return
		}

		respondWithJson(w, http.StatusOK, chirps)
		return
	}

	chirps, err := cfg.db.ReadAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retrieving chirps from database", err)
		return
	}

	if q := r.URL.Query().Get("sort"); q == "desc" {
		sort.Slice(chirps, func(i, j int) bool {
			if chirps[i].CreatedAt.Compare(chirps[j].CreatedAt) == +1 {
				return true
			} else {
				return false
			}
		})
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
