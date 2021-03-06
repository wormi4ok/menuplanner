package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/wormi4ok/menuplanner/internal"
)

type courseEndpoint struct {
	storage internal.CourseReader
}

// Routes creates a REST router for the course resource
func (e courseEndpoint) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", e.List())
	r.Get("/{id}", e.Get())

	return r
}

func (e courseEndpoint) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseJSON(w, e.storage.ReadAllCourses(r.Context()))
	}
}

func (e courseEndpoint) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		course := e.storage.ReadCourse(r.Context(), id)
		if course == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		responseJSON(w, course)
	}
}
