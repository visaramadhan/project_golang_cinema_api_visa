package routes

import (
	"github.com/bioskop-api/api/cinemas"
	"github.com/go-chi/chi/v5"
)

func CinemaRoutes(router chi.Router, cinemaHandler *cinemas.CinemaHandler) {
	router.Get("/api/cinemas", cinemaHandler.GetCinemas)
	router.Get("/api/cinemas/{id}", cinemaHandler.GetCinemaByID)
}
