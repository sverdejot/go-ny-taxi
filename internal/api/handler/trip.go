package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sverdejot/go-ny-taxi/internal/model"
)

type TripHandler struct {
	Repo model.TripRepository
}

func (h *TripHandler) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		trip, found := h.Repo.Find(id)

		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resp, err := json.Marshal(trip)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(resp))
	})
}

func (h *TripHandler) Post() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var trip model.Trip
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)

		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&trip); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.Repo.Add(trip); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}
