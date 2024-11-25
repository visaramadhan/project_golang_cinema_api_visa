package paymentmethod

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bioskop-api/utils/database"
)

func GetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	database.InitDB()
	db := database.DB
	repo := PaymentMethodRepository{DB: db}
	methods, err := repo.GetAllPaymentMethods()
	if err != nil {
		fmt.Println("err:", err)
		http.Error(w, "Failed to fetch payment methods", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"statusCode": 200,
		"data":       methods,
	})
}
