package api

import (
	"log"
	"net/http"
	"time"

	"github.com/sverdejot/go-ny-taxi/internal/api/handler"
	"github.com/sverdejot/go-ny-taxi/internal/storage/postgres"
)

func Run() {
	db := storage.Initialize()
	repo := storage.NewPostgresTripRepository(db)

	handler := handler.TripHandler{
		Repo: repo,
	}

	mux := http.NewServeMux()
	mux.Handle("GET /trips/{id}", handler.Get())
	mux.Handle("POST /trips", handler.Post())

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
