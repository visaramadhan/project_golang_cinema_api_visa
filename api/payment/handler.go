package payment

import (
	"encoding/json"
	"net/http"

	"github.com/bioskop-api/utils/database"
)

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var paymentRequest Payment
	if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if paymentRequest.PaymentMethod == "" || paymentRequest.BookingID == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	database.InitDB()
	db := database.DB
	repo := PaymentRepository{DB: db}
	transactionID, err := repo.ProcessPayment(&paymentRequest)
	if err != nil {
		http.Error(w, "Payment failed", http.StatusPaymentRequired)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"statusCode":    200,
		"message":       "Payment successful.",
		"transactionId": transactionID,
		"bookingId":     paymentRequest.BookingID,
	})
}
