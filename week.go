package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/wormi4ok/menuplanner/internal"
)

type weekEndpoint struct {
	storage internal.WeekRepository
}

// Routes creates a REST router for the todos resource
func (e weekEndpoint) Routes() chi.Router {
	r := chi.NewRouter()

	r.Put("/", e.Update())
	r.Get("/", e.Get())

	return r
}

func (e *weekEndpoint) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(e.storage.ReadCurrent())
		if err != nil {
			log.Printf("Handler error: %v", err)
			w.WriteHeader(500)
		}
	}
}

func (e *weekEndpoint) Update() http.HandlerFunc {
	type request struct {
		internal.Week
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			_, _ = io.WriteString(w, "Missing or malformed payload")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := json.Unmarshal(body, req); err != nil {
			_, _ = io.WriteString(w, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		e.storage.UpdateCurrent(&req.Week)
	}
}
