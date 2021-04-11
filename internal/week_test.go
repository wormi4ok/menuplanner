package internal

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type mock struct {
	Recipes []*Recipe
}

func (m *mock) Read(_ context.Context, id int) *Recipe {
	for _, r := range m.Recipes {
		if r.ID == id {
			return r
		}
	}
	return nil
}

func (m *mock) ReadAll(_ context.Context) []*Recipe {
	return m.Recipes
}

func TestGapFiller_FillWeek(t *testing.T) {
	tests := []struct {
		name    string
		recipes []*Recipe
		input   *Week
		want    *Week
	}{
		{
			name:    "Happy path",
			recipes: recipesFromJSON("recipes.json"),
			input:   weekFromJSON("week_with_gaps.json"),
			want:    weekFromJSON("week_golden.json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gf := &GapFiller{
				r: &mock{tt.recipes},
			}
			got := gf.FillWeek(context.TODO(), tt.input)
			for i, day := range tt.want.Menu {
				if _, exists := got.Menu[i]; !exists {
					t.Errorf("Missing menu for day %d", i)
					continue
				}
				gotMenu := got.Menu[i]
				for j, recipe := range day.Recipes {
					if _, exists := gotMenu.Recipes[j]; !exists {
						t.Errorf("Missing recipe for day %d, slot %d", i, j)
						continue
					}
					if recipe.ID != got.Menu[i].Recipes[j].ID {
						t.Errorf("Recipe on day %d, slot %d didn't match expected: want = %d, got = %d", i, j, recipe.ID, gotMenu.Recipes[j].ID)
					}
				}
			}
		})
	}
}

func recipesFromJSON(file string) []*Recipe {
	path := filepath.Join("testdata", file)
	data, err := os.ReadFile(path)
	if err != nil {
		panic("failed to load file")
	}
	var rr []*Recipe

	if err = json.Unmarshal(data, &rr); err != nil {
		panic("failed to parse JSON")
	}

	return rr
}

func weekFromJSON(file string) *Week {
	path := filepath.Join("testdata", file)
	data, err := os.ReadFile(path)
	if err != nil {
		panic("failed to load file")
	}
	var w *Week

	if err = json.Unmarshal(data, &w); err != nil {
		panic("failed to parse JSON")
	}
	return w
}
