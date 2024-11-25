package seats

import (
	"database/sql"
	"errors"
)

type SeatRepository struct {
	DB *sql.DB
}

func NewSeatRepository(db *sql.DB) *SeatRepository {
	return &SeatRepository{
		DB: db,
	}
}

func (sr *SeatRepository) GetSeatsByCinema(cinemaID int) ([]Seat, error) {
	query := "SELECT id, cinema_id, seat_number, is_booked FROM seats WHERE cinema_id = $1"
	rows, err := sr.DB.Query(query, cinemaID)
	if err != nil {
		return nil, errors.New("failed to fetch seats: " + err.Error())
	}
	defer rows.Close()

	var seats []Seat
	for rows.Next() {
		var seat Seat
		err := rows.Scan(&seat.ID, &seat.CinemaID, &seat.SeatNumber, &seat.IsBooked)
		if err != nil {
			return nil, errors.New("error scanning seat row: " + err.Error())
		}
		seats = append(seats, seat)
	}
	return seats, nil
}
