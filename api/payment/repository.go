package payment

import (
	"database/sql"
	"fmt"
)

type PaymentRepository struct {
	DB *sql.DB
}

func (r *PaymentRepository) ProcessPayment(payment *Payment) (string, error) {

	transactionID := fmt.Sprintf("txn%s", payment.BookingID)
	_, err := r.DB.Exec(`
		INSERT INTO payments (booking_id, payment_method, card_number, expiry_date, cvv, transaction_id)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		payment.BookingID, payment.PaymentMethod, payment.CardNumber, payment.ExpiryDate, payment.CVV, transactionID)

	if err != nil {
		return "", err
	}
	return transactionID, nil
}
