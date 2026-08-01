package internal

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

const testUserID = 0

var slotCourse = map[int]string{0: CourseBreakfast, 1: CourseMain, 2: CourseMain}

func TestGapFiller_FillWeek(t *testing.T) {
	m := &mock{Recipes: recipesFromJSON("recipes.json"), Courses: coursesFromJSON("courses.json")}
	gf := &GapFiller{r: m, c: m}
	input := weekFromJSON("week_with_gaps.json")

	preset := map[[2]int]int{}
	for i, day := range input.Menu {
		for j, recipe := range day.Recipes {
			preset[[2]int{i, j}] = recipe.ID
		}
	}

	got := gf.FillWeek(context.TODO(), testUserID, input)

	for i := 0; i < 7; i++ {
		day, exists := got.Menu[i]
		if !exists {
			t.Errorf("Missing menu for day %d", i)
			continue
		}
		for j := 0; j < 3; j++ {
			recipe, exists := day.Recipes[j]
			if !exists || recipe.IsEmpty() {
				t.Errorf("Day %d, slot %d was left empty", i, j)
				continue
			}
			if id, ok := preset[[2]int{i, j}]; ok && id != recipe.ID {
				t.Errorf("Day %d, slot %d was already filled with recipe %d, got %d", i, j, id, recipe.ID)
			}
			if !servesCourse(recipe, slotCourse[j]) {
				t.Errorf("Day %d, slot %d wants a %s, got recipe %d which is not one", i, j, slotCourse[j], recipe.ID)
			}
		}
	}
}

func servesCourse(r Recipe, name string) bool {
	for _, c := range r.Courses {
		if c.Name == name {
			return true
		}
	}
	return false
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
