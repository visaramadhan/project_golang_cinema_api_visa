package seats

type Seat struct {
	ID         int    `json:"id"`
	CinemaID   int    `json:"cinema_id"`
	SeatNumber string `json:"seat_number"`
	IsBooked   bool   `json:"is_booked"`
}
