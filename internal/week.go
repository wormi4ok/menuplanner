package internal

import (
	"context"
	"errors"
)

type WeekRepository interface {
	ReadCurrent(context.Context) *Week
	UpdateCurrent(context.Context, *Week) *Week
	DeleteSlot(ctx context.Context, week, day, slot int) error
}

type DailyMenu struct {
	Recipes  map[int]Recipe `json:"recipes"`
	calories int
}

type Week struct {
	Menu map[int]*DailyMenu `json:"menu"`
}

type GapFiller struct {
	r RecipeReader
	c CourseReader

	indexedCourses map[int]*Course
}

func NewGapFiller(recipes RecipeReader, courses CourseReader) *GapFiller {
	return &GapFiller{r: recipes, c: courses}
}

func (gf *GapFiller) FillWeek(ctx context.Context, week *Week) *Week {
	gf.prepareSkeleton(ctx, week)
	for i, menu := range week.Menu {
		for j, recipe := range menu.Recipes {
			if recipe.IsEmpty() {
				for attempts := 3; attempts > 0; attempts-- {
					r := *gf.r.ReadRandom(ctx, *gf.courseName(j))
					week.Menu[i].AddRecipe(j, r)
					if menu.CheckRecipe(r) == nil {
						break
					}
				}
			}
		}
	}

	return week
}
func (gf *GapFiller) courseName(index int) *Course {
	if gf.indexedCourses != nil {
		return gf.indexedCourses[index]
	}

	cc := gf.c.ReadAllCourses(context.Background())
	gf.indexedCourses = make(map[int]*Course, 3)
	for _, course := range cc {
		switch course.Name {
		case CourseBreakfast:
			gf.indexedCourses[0] = course // breakfast
		case CourseMain:
			gf.indexedCourses[1] = course // lunch
			gf.indexedCourses[2] = course // dinner
		case CoursePudding:
			gf.indexedCourses[3] = course // pudding
		}
	}

	return gf.indexedCourses[index]
}

func (gf *GapFiller) prepareSkeleton(ctx context.Context, week *Week) {
	for i := 0; i < 7; i++ {
		if _, exists := week.Menu[i]; !exists {
			week.Menu[i] = &DailyMenu{Recipes: map[int]Recipe{}}
		}
		for j := 0; j < 3; j++ {
			if _, exists := week.Menu[i].Recipes[j]; !exists {
				week.Menu[i].Recipes[j] = Recipe{}
			}
			if !week.Menu[i].Recipes[j].IsEmpty() {
				r := gf.r.Read(ctx, week.Menu[i].Recipes[j].ID)
				if r != nil {
					week.Menu[i].AddRecipe(j, *r)
				}
			}
		}
	}
}

func (dm *DailyMenu) AddRecipe(slot int, recipe Recipe) {
	if r, exists := dm.Recipes[slot]; exists && !r.IsEmpty() {
		dm.calories -= r.EnergyAmount()
	}

	dm.calories += recipe.EnergyAmount()
	dm.Recipes[slot] = recipe
}

func (dm *DailyMenu) CheckRecipe(recipe Recipe) error {
	if dm.calories+recipe.EnergyAmount() > MaxCalories {
		return errors.New("max calories per day exceeded")
	}
	for _, r := range dm.Recipes {
		if r.ID == recipe.ID {
			return errors.New("duplicate recipe")
		}
	}

	return nil
}
