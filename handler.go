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
)

type recipeResource struct {
	validate *validator.Validate
}

// Routes creates a REST router for the todos resource
func (rs recipeResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)
	r.Post("/", rs.Create)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(rs.RecipeCtx)
		r.Get("/", rs.Get)
		r.Delete("/", rs.Delete)
	})

	return r
}

func (rs recipeResource) RecipeCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		var recipe *Recipe
		for _, r := range allRecipes {
			if r.ID == id {
				recipe = &r
				break
			}
		}
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

func (rs recipeResource) Create(w http.ResponseWriter, r *http.Request) {
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

	if err := rs.validate.StructCtx(r.Context(), req); err != nil {
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

	allRecipes = append(allRecipes, Recipe{
		ID:          len(allRecipes) + 1,
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

func (rs recipeResource) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(allRecipes)
	if err != nil {
		log.Printf("Handler error: %v", err)
		w.WriteHeader(500)
	}
}

func (rs recipeResource) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	recipe, ok := ctx.Value("recipe").(*Recipe)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	for i, r := range allRecipes {
		if r.ID == recipe.ID {
			allRecipes[i] = allRecipes[len(allRecipes)-1]
			allRecipes = allRecipes[:len(allRecipes)-1]
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

func (rs recipeResource) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	recipe, ok := ctx.Value("recipe").(*Recipe)
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
