package internal

import "context"

type RecipeRepository interface {
	RecipeReader

	Create(ctx context.Context, r *Recipe) *Recipe
	Delete(ctx context.Context, id int) bool
}

type RecipeReader interface {
	Read(ctx context.Context, id int) *Recipe
	ReadAll(ctx context.Context) []*Recipe
}

type Recipe struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"imageUrl"`

	Calories int `json:"calories"`
	Protein  int `json:"protein"`
	Fat      int `json:"fat"`
	Carbs    int `json:"carbs"`
}
