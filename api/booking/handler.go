package booking

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bioskop-api/utils/database"
	"github.com/google/uuid"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var bookingRequest Booking
	if err := json.NewDecoder(r.Body).Decode(&bookingRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	database.InitDB()
	db := database.DB
	repo := BookingRepository{DB: db}
	isBooked, err := repo.CheckSeatAvailability(bookingRequest.SeatID)
	if err != nil {
		http.Error(w, "Error checking seat availability", http.StatusInternalServerError)
		return
	}

	if isBooked {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": 409,
			"message":    "The selected seat is already booked.",
		})
		return
	}

	bookingRequest.BookingID = uuid.New().String()
	err = repo.CreateBooking(&bookingRequest)
	if err != nil {
		http.Error(w, "Failed to create booking", http.StatusInternalServerError)
		log.Println("Error creating booking:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"statusCode": 201,
		"message":    "Booking confirmed.",
		"booking":    bookingRequest,
	})
}
