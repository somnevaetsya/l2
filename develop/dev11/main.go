package main

import (
	"l2/develop/dev11/app/handler"
	"l2/develop/dev11/app/repository"
	"l2/develop/dev11/app/usecase"
	"l2/develop/dev11/pkg/middleware"
	"log"
	"net/http"
)

func main() {
	repo := repository.MakeRepository()
	use := usecase.MakeUsecase(repo)
	hand := handler.MakeHandler(use)
	router := http.NewServeMux()
	router.Handle("/create_event", middleware.LogMiddleware(http.HandlerFunc(hand.CreateEvent)))
	router.Handle("/update_event", middleware.LogMiddleware(http.HandlerFunc(hand.UpdateEvent)))
	router.Handle("/delete_event", middleware.LogMiddleware(http.HandlerFunc(hand.DeleteEvent)))
	router.Handle("/events_for_day", middleware.LogMiddleware(http.HandlerFunc(hand.GetDayEvents)))
	router.Handle("/events_for_week", middleware.LogMiddleware(http.HandlerFunc(hand.GetWeekEvents)))
	router.Handle("/events_for_month", middleware.LogMiddleware(http.HandlerFunc(hand.GetMonthEvents)))

	server := &http.Server{
		Addr:    "localhost:5555",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server error: %s", err)
	}
}
