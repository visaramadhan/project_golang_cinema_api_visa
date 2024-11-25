package main

import (
	"log"
	"net/http"

	"github.com/bioskop-api/api/booking"
	"github.com/bioskop-api/api/cinemas"
	"github.com/bioskop-api/api/payment"
	paymentmethod "github.com/bioskop-api/api/payment_method"
	"github.com/bioskop-api/api/seats"
	"github.com/bioskop-api/api/users"
	"github.com/bioskop-api/utils/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	database.InitDB()
	db := database.DB

	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo)
	userHandler := users.NewUserHandler(userService)

	cinemaRepo := cinemas.NewCinemaRepository(db)
	cinemaHandler := cinemas.NewCinemaHandler(cinemaRepo)

	seatRepo := seats.NewSeatRepository(db)
	seatHandler := seats.NewSeatHandler(seatRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/api/register", userHandler.Register)
	r.Post("/api/login", userHandler.Login)

	r.Get("/api/cinemas", cinemaHandler.GetCinemas)
	r.Get("/api/cinemas/{cinemaId}", cinemaHandler.GetCinemaByID)

	r.Get("/api/seats", seatHandler.GetSeatsByCinema)

	r.Post("/api/pay", payment.ProcessPayment)

	r.Get("/api/payment-methods", paymentmethod.GetPaymentMethods)

	r.Post("/api/booking", booking.CreateBooking)

	// Start server
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
