package payment

type Payment struct {
	BookingID     string `json:"bookingId"`
	PaymentMethod string `json:"paymentMethod"`
	CardNumber    string `json:"cardNumber,omitempty"`
	ExpiryDate    string `json:"expiryDate,omitempty"`
	CVV           string `json:"cvv,omitempty"`
	TransactionID string `json:"transactionId"`
}
