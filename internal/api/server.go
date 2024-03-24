package api

import (
	"log"
	"net/http"
	"time"

	"github.com/sverdejot/go-ny-taxi/internal/api/handler"
	"github.com/sverdejot/go-ny-taxi/internal/storage"
)

func Run() {
	mux := http.NewServeMux()

	repo := storage.NewInMemoryTripRepository()

	handler := handler.TripHandler{
		Repo: repo,
	}

	mux.Handle("GET /trip/{id}", handler.Get())

	srv := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,

		Handler: mux,
	}

	log.Println("starting sever")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
