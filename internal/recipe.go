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

const (
	CourseBreakfast = "breakfast"
	CourseMain      = "main"
	CoursePudding   = "pudding"
)

type Course struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"not null;size:255"`
}

type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"name" validate:"required" gorm:"not null"`
	Courses     []Course `json:"courses" gorm:"many2many:recipe_courses;"`
	Description string   `json:"description,omitempty"`
	ImageURL    string   `json:"imageUrl" validate:"omitempty,url"`

	Calories int `json:"calories" validate:"required"`
	Protein  int `json:"protein" validate:"required"`
	Fat      int `json:"fat" validate:"required"`
	Carbs    int `json:"carbs" validate:"required"`

	Quantity int `json:"quantity,omitempty"`
	Portion  int `json:"portion,omitempty" validate:"required_with=Quantity"`
}

func (r *Recipe) EnergyAmount() int {
	if r.Quantity > 0 {
		return r.Calories / r.Quantity * r.Portion
	}

	return r.Calories
}

type Validator interface {
	StructCtx(ctx context.Context, s interface{}) (err error)
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
