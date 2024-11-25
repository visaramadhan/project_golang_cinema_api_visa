package booking

type Booking struct {
	BookingID     string `json:"bookingId"`
	CinemaID      int    `json:"cinemaId"`
	SeatID        string `json:"seatId"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	PaymentMethod string `json:"paymentMethod"`
}
