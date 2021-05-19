package internal

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	jwt2 "github.com/wormi4ok/menuplanner/internal/http/jwt"
)

func TestGapFiller_FillWeek(t *testing.T) {
	tests := []struct {
		name    string
		recipes []*Recipe
		courses []*Course
		input   *Week
		want    *Week
	}{
		{
			name:    "Happy path",
			recipes: recipesFromJSON("recipes.json"),
			courses: coursesFromJSON("courses.json"),
			input:   weekFromJSON("week_with_gaps.json"),
			want:    weekFromJSON("week_golden.json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mock{Recipes: tt.recipes, Courses: tt.courses}
			gf := &GapFiller{
				r: m,
				c: m,
			}
			got := gf.FillWeek(context.TODO(), jwt2.UserID(r.Context()), tt.input)
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

func recipesFromJSON(file string) (rr []*Recipe) {
	path := filepath.Join("testdata", file)
	data, err := os.ReadFile(path)
	if err != nil {
		panic("failed to load file")
	}

	if err = json.Unmarshal(data, &rr); err != nil {
		panic("failed to parse JSON")
	}

	return rr
}

func coursesFromJSON(file string) (cc []*Course) {
	path := filepath.Join("testdata", file)
	data, err := os.ReadFile(path)
	if err != nil {
		panic("failed to load file")
	}

	if err = json.Unmarshal(data, &cc); err != nil {
		panic("failed to parse JSON")
	}

	return cc
}

func weekFromJSON(file string) (w *Week) {
	path := filepath.Join("testdata", file)
	data, err := os.ReadFile(path)
	if err != nil {
		panic("failed to load file")
	}

	if err = json.Unmarshal(data, &w); err != nil {
		panic("failed to parse JSON")
	}
	return w
}

type mock struct {
	Recipes []*Recipe
	Courses []*Course

	seed int
}

func (m *mock) Read(ctx context.Context, userID int, id int) *Recipe {
	for _, r := range m.Recipes {
		if r.ID == id {
			return r
		}
	}
	return nil
}

func (m *mock) ReadAll(ctx context.Context, userID int) []*Recipe {
	return m.Recipes
}

func (m *mock) ReadRandom(ctx context.Context, course Course, userID int) *Recipe {
	var rr []*Recipe
	for i := 0; i < len(m.Recipes); i++ {
		r := m.Recipes[i]
		for j := 0; j < len(r.Courses); j++ {
			c := r.Courses[j]
			if course.ID == c.ID {
				rr = append(rr, r)
			}
		}
	}

	if rr == nil {
		return nil
	}
	m.seed++
	if m.seed >= len(rr) {
		m.seed = m.seed - len(rr)
	}
	return rr[m.seed]
}

func (m *mock) ReadCourse(_ context.Context, id int) *Course {
	for _, course := range m.Courses {
		if course.ID != id {
			return course
		}
	}
	return nil
}

func (m *mock) ReadAllCourses(_ context.Context) []*Course {
	return m.Courses
}
