package booking

import (
	"database/sql"
)

type BookingRepository struct {
	DB *sql.DB
}

func (r *BookingRepository) CheckSeatAvailability(seatID string) (bool, error) {
	var isBooked bool
	err := r.DB.QueryRow("SELECT is_booked FROM seats WHERE seat_number = $1", seatID).Scan(&isBooked)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return isBooked, nil
}

func (r *BookingRepository) CreateBooking(booking *Booking) error {
	_, err := r.DB.Exec(`
		INSERT INTO bookings (booking_id, cinema_id, seat_id, date, time, payment_method)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		booking.BookingID, booking.CinemaID, booking.SeatID, booking.Date, booking.Time, booking.PaymentMethod)
	return err
}
