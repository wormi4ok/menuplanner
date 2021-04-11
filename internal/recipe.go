package internal

import (
	"context"

	"github.com/go-playground/validator/v10"
)

type RecipeRepository interface {
	RecipeWriter
	RecipeReader
}

type RecipeReader interface {
	Read(ctx context.Context, id int) *Recipe
	ReadAll(ctx context.Context) []*Recipe
}

type RecipeWriter interface {
	Create(ctx context.Context, r *Recipe) (*Recipe, error)
	Delete(ctx context.Context, id int) bool
}

type Validator interface {
	StructCtx(ctx context.Context, s interface{}) (err error)
}

type Recipe struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"imageUrl" validate:"omitempty,url"`

	Calories int `json:"calories" validate:"required"`
	Protein  int `json:"protein" validate:"required"`
	Fat      int `json:"fat" validate:"required"`
	Carbs    int `json:"carbs" validate:"required"`
}

func SaveRecipe(ctx context.Context, recipe Recipe, storage RecipeWriter) (id int, err error) {
	v := validator.New()

	err = v.StructCtx(ctx, recipe)
	if err != nil {
		return
	}

	if r, err := storage.Create(ctx, &recipe); err == nil {
		id = r.ID
	}
	return
}
