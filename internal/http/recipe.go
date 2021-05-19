package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/wormi4ok/menuplanner/internal"
	"github.com/wormi4ok/menuplanner/internal/http/jwt"
)

type recipeEndpoint struct {
	storage internal.RecipeRepository
}

// Routes creates a REST router for the recipe resource
func (e recipeEndpoint) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", e.List())
	r.Post("/", e.Create())

	r.Route("/{id}", func(r chi.Router) {
		r.Use(e.RecipeCtx)
		r.Get("/", e.Get())
		r.Put("/", e.Update())
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
		recipe := e.storage.Read(r.Context(), jwt.UserID(r.Context()), id)
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
		var req request

		if err := readJSON(r, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := internal.SaveRecipe(r.Context(), jwt.UserID(r.Context()), internal.Recipe(req), e.storage)
		if err != nil {
			if internal.ErrorIs(err, internal.ErrorUnauthorized) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

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
		req.ID = id

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(req); err != nil {
			_, _ = io.WriteString(w, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (e recipeEndpoint) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		responseJSON(w, e.storage.ReadAll(ctx, jwt.UserID(ctx)))
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

		if err := internal.DeleteRecipe(ctx, jwt.UserID(r.Context()), recipe, e.storage); err != nil {
			if internal.ErrorIs(err, internal.ErrorUnauthorized) {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
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

		responseJSON(w, recipe)
	}
}

func (e recipeEndpoint) Update() http.HandlerFunc {
	type request internal.Recipe

	type validationError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	type validationErrorResponse struct {
		Errors []validationError `json:"errors"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req    request
			ctx    = r.Context()
			userID = jwt.UserID(r.Context())
		)

		if _, ok := ctx.Value("recipe").(*internal.Recipe); !ok {
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}

		if err := readJSON(r, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		recipe := internal.Recipe(req)

		_, err := internal.UpdateRecipe(r.Context(), userID, &recipe, e.storage)
		if err != nil {
			if internal.ErrorIs(err, internal.ErrorUnauthorized) {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
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

		w.WriteHeader(http.StatusAccepted)
		responseJSON(w, recipe)
	}
}
