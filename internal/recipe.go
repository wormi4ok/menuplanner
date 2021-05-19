package internal

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
)

type RecipeRepository interface {
	RecipeWriter
	RecipeReader
}

type RecipeReader interface {
	Read(ctx context.Context, userID int, id int) *Recipe
	ReadAll(ctx context.Context, userID int) []*Recipe
	ReadRandom(ctx context.Context, course Course, userID int) *Recipe
}

type RecipeWriter interface {
	Create(ctx context.Context, r *Recipe) (*Recipe, error)
	Update(ctx context.Context, r *Recipe) (*Recipe, error)
	Delete(ctx context.Context, id int) bool
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

	UserID int  `json:"-" validate:"required" gorm:"not null;default:1"`
	User   User `json:"-"`
}

func (r Recipe) EnergyAmount() int {
	if r.Quantity > 0 {
		return r.Calories / r.Quantity * r.Portion
	}

	return r.Calories
}

func (r Recipe) IsEmpty() bool {
	return r.ID == 0
}

type Validator interface {
	StructCtx(ctx context.Context, s interface{}) (err error)
}

func SaveRecipe(ctx context.Context, userID int, recipe Recipe, storage RecipeWriter) (id int, err error) {
	recipe.ensureUserID(userID)
	v := validator.New()

	err = v.StructCtx(ctx, recipe)
	if err != nil {
		return
	}

	if recipe.UserID != userID {
		err = NewError(errors.New("user ID doesn't match recipe.User"), ErrorUnauthorized)
		return
	}

	if r, err := storage.Create(ctx, &recipe); err == nil {
		id = r.ID
	}
	return
}

func UpdateRecipe(ctx context.Context, userID int, r *Recipe, storage RecipeWriter) (recipe *Recipe, err error) {
	r.ensureUserID(userID)
	v := validator.New()

	err = v.StructCtx(ctx, r)
	if err != nil {
		return
	}

	if r.UserID != userID {
		err = NewError(errors.New("user is not authorized to modify this recipe"), ErrorUnauthorized)
		return
	}

	recipe, err = storage.Update(ctx, r)
	return
}

func DeleteRecipe(ctx context.Context, userID int, recipe *Recipe, storage RecipeWriter) error {
	if userID != recipe.UserID {
		return NewError(errors.New("user is not authorized to modify this recipe"), ErrorUnauthorized)
	}

	if !storage.Delete(ctx, recipe.ID) {
		return errors.New("failed to delete database record")
	}

	return nil
}

func (r *Recipe) ensureUserID(userID int) {
	if r.UserID == 0 {
		r.UserID = userID
	}
}
