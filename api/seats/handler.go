package seats

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type SeatHandler struct {
	Repo *SeatRepository
}

func NewSeatHandler(repo *SeatRepository) *SeatHandler {
	return &SeatHandler{Repo: repo}
}

func (sh *SeatHandler) GetSeatsByCinema(w http.ResponseWriter, r *http.Request) {
	cinemaIDStr := r.URL.Query().Get("cinema_id")
	cinemaID, err := strconv.Atoi(cinemaIDStr)
	if err != nil {
		fmt.Println("err:", err)
		http.Error(w, "Invalid cinema ID", http.StatusBadRequest)
		return
	}

	seats, err := sh.Repo.GetSeatsByCinema(cinemaID)
	if err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "Failed to fetch seats: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"statusCode": http.StatusOK,
		"data":       seats,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
