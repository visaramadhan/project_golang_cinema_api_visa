package cinemas

import (
	"database/sql"
	"errors"
)

type CinemaRepository struct {
	DB *sql.DB
}

func NewCinemaRepository(db *sql.DB) *CinemaRepository {
	return &CinemaRepository{
		DB: db,
	}
}

func (cr *CinemaRepository) GetCinemas(limit, offset int) ([]Cinema, error) {
	query := "SELECT id, name, location, seats FROM cinemas LIMIT $1 OFFSET $2"
	rows, err := cr.DB.Query(query, limit, offset)
	if err != nil {
		return nil, errors.New("failed to fetch cinemas: " + err.Error())
	}
	defer rows.Close()

	var cinemas []Cinema
	for rows.Next() {
		var cinema Cinema
		err := rows.Scan(&cinema.ID, &cinema.Name, &cinema.Location, &cinema.Seats)
		if err != nil {
			return nil, errors.New("error scanning cinema row: " + err.Error())
		}
		cinemas = append(cinemas, cinema)
	}
	return cinemas, nil
}

func (cr *CinemaRepository) GetCinemaByID(id int) (Cinema, error) {
	query := "SELECT id, name, location, seats FROM cinemas WHERE id = $1"
	row := cr.DB.QueryRow(query, id)

	var cinema Cinema
	err := row.Scan(&cinema.ID, &cinema.Name, &cinema.Location, &cinema.Seats)
	if err == sql.ErrNoRows {
		return Cinema{}, errors.New("cinema not found")
	} else if err != nil {
		return Cinema{}, errors.New("error retrieving cinema: " + err.Error())
	}
	return cinema, nil
}
