package main

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
	storage  internal.RecipeRepository
	validate *validator.Validate
}

// Routes creates a REST router for the todos resource
func (e recipeEndpoint) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", e.List)
	r.Post("/", e.Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(e.RecipeCtx)
		r.Get("/", e.Get)
		r.Delete("/", e.Delete)
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

type AddRecipeRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl" validate:"omitempty,url"`

	Calories int `json:"calories" validate:"required"`
	Protein  int `json:"protein" validate:"required"`
	Fat      int `json:"fat" validate:"required"`
	Carbs    int `json:"carbs" validate:"required"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Errors []ValidationError `json:"errors"`
}

func (e recipeEndpoint) Create(w http.ResponseWriter, r *http.Request) {
	req := &AddRecipeRequest{}

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

	if err := e.validate.StructCtx(r.Context(), req); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			_, _ = io.WriteString(w, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res := ValidationErrorResponse{}
		for _, err := range err.(validator.ValidationErrors) {

			res.Errors = append(res.Errors, ValidationError{
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

	e.storage.Create(r.Context(), &internal.Recipe{
		Name:        req.Name,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		Calories:    req.Calories,
		Protein:     req.Protein,
		Fat:         req.Fat,
		Carbs:       req.Carbs,
	})

	w.WriteHeader(http.StatusCreated)
}

func (e recipeEndpoint) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(e.storage.ReadAll(r.Context()))
	if err != nil {
		log.Printf("Handler error: %v", err)
		w.WriteHeader(500)
	}
}

func (e recipeEndpoint) Delete(w http.ResponseWriter, r *http.Request) {
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

func (e recipeEndpoint) Get(w http.ResponseWriter, r *http.Request) {
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
