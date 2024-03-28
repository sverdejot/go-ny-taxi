package api

import (
	"log"
	"net/http"
	"time"

	"github.com/sverdejot/go-ny-taxi/internal/api/handler"
	"github.com/sverdejot/go-ny-taxi/internal/config"
	"github.com/sverdejot/go-ny-taxi/internal/storage/postgres"
)

func Run() {
	config, err := config.Init()
	if err != nil {
		log.Fatal("could not load config")
	}

	cs := storage.NewConnectionString(
		storage.WithDriver("postgres"),
		storage.WithHost(config.Database.Host),
		storage.WithPort(config.Database.Port),
		storage.WithUser(config.Database.User),
		storage.WithPassword(config.Database.Password),
		storage.WithDatabase(config.Database.Database),
		storage.WithOpts(map[string]string{
			"sslmode": "disable",
		}),
	)

	db := storage.Init(cs.String())
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
