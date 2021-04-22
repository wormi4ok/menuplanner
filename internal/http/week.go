package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/wormi4ok/menuplanner/internal"
)

type weekEndpoint struct {
	storage internal.WeekRepository
	filler  *internal.GapFiller
}

// Routes creates a REST router for the week resource
func (e weekEndpoint) Routes() chi.Router {
	r := chi.NewRouter()

	r.Put("/", e.Update())
	r.Get("/", e.Get())
	r.Delete("/day/{day}/slot/{slot}", e.Delete())

	return r
}

func (e *weekEndpoint) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(e.storage.ReadCurrent(r.Context()))
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

		week := &req.Week
		if r.URL.Query().Get("fillGaps") != "" {
			e.filler.FillWeek(r.Context(), week)
		}
		res := e.storage.UpdateCurrent(r.Context(), week)

		w.WriteHeader(http.StatusAccepted)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			_, _ = io.WriteString(w, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (e *weekEndpoint) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		day, err := strconv.Atoi(chi.URLParam(r, "day"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		slot, err := strconv.Atoi(chi.URLParam(r, "slot"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		err = e.storage.DeleteSlot(r.Context(), 1, day, slot)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
