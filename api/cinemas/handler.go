package cinemas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type CinemaHandler struct {
	Repo *CinemaRepository
}

func NewCinemaHandler(repo *CinemaRepository) *CinemaHandler {
	return &CinemaHandler{Repo: repo}
}

func (ch *CinemaHandler) GetCinemas(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	cinemas, err := ch.Repo.GetCinemas(limit, offset)
	if err != nil {
		http.Error(w, "Failed to fetch cinemas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"statusCode": http.StatusOK,
		"data":       cinemas,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *CinemaHandler) GetCinemaByID(w http.ResponseWriter, r *http.Request) {
	cinemaID := chi.URLParam(r, "cinemaId")

	if cinemaID == "" {
		http.Error(w, "cinemaId cannot be empty", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(cinemaID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid cinema ID: %v", err), http.StatusBadRequest)
		return
	}

	cinema, err := h.Repo.GetCinemaByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"statusCode": 200,
		"data":       cinema,
	})
}
