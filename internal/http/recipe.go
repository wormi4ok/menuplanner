package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/wormi4ok/menuplanner/internal"
)

type recipeEndpoint struct {
	storage internal.RecipeRepository
}

// Routes creates a REST router for the todos resource
func (e recipeEndpoint) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", e.List())
	r.Post("/", e.Create())

	r.Route("/{id}", func(r chi.Router) {
		r.Use(e.RecipeCtx)
		r.Get("/", e.Get())
		r.Delete("/", e.Delete())
	})

	return r
}

func (e recipeEndpoint) RecipeCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		recipe := e.storage.Read(r.Context(), id)
		if recipe == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "recipe", recipe)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (e recipeEndpoint) Create() http.HandlerFunc {
	type request internal.Recipe

	type validationError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	type validationErrorResponse struct {
		Errors []validationError `json:"errors"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			_, _ = io.WriteString(w, "Missing or malformed payload")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := json.Unmarshal(body, &req); err != nil {
			_, _ = io.WriteString(w, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if _, err := internal.SaveRecipe(r.Context(), internal.Recipe(req), e.storage); err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				_, _ = io.WriteString(w, err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			res := validationErrorResponse{}
			for _, err := range err.(validator.ValidationErrors) {

				res.Errors = append(res.Errors, validationError{
					Field:   err.Field(),
					Message: fmt.Sprintf("Validation error (%s)", err.Tag()),
				})
			}

			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(res); err != nil {
				_, _ = io.WriteString(w, err.Error())
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (e recipeEndpoint) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(e.storage.ReadAll(r.Context()))
		if err != nil {
			log.Printf("Handler error: %v", err)
			w.WriteHeader(500)
		}
	}
}

func (e recipeEndpoint) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		recipe, ok := ctx.Value("recipe").(*internal.Recipe)
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		if e.storage.Delete(r.Context(), recipe.ID) {
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

func (e recipeEndpoint) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		recipe, ok := ctx.Value("recipe").(*internal.Recipe)
		if !ok {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(recipe)
		if err != nil {
			log.Printf("Handler error: %v", err)
			w.WriteHeader(500)
		}
	}
}
