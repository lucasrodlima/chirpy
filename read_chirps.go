package main

import "net/http"

func (cfg *apiConfig) handlerReadChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.ReadAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retrieving chirps from database", err)
		return
	}

	respondWithJson(w, http.StatusOK, chirps)
}
