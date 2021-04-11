package internal

import "context"

type WeekRepository interface {
	ReadCurrent() *Week
	UpdateCurrent(*Week) *Week
}

type DailyMenu struct {
	Recipes map[int]*Recipe `json:"recipes"`
}

type Menu map[int]DailyMenu

type Week struct {
	Menu Menu `json:"menu"`
}

type GapFiller struct {
	r RecipeReader
}

func NewGapFiller(recipes RecipeReader) *GapFiller {
	return &GapFiller{r: recipes}
}

func (gf *GapFiller) FillWeek(ctx context.Context, week *Week) *Week {
	for i := 0; i < 2; i++ {
		if _, exists := week.Menu[i]; !exists {
			week.Menu[i] = DailyMenu{}
		}

		todayRecipeIDs := make(map[int]*Recipe, 3)
		dayCalories := 0

		for j := 0; j < 3; j++ {
			_, exists := week.Menu[i].Recipes[j]
			if !exists {
				week.Menu[i].Recipes[j] = &Recipe{}
			}

			id := week.Menu[i].Recipes[j].ID
			if id == 0 {
				for _, recipe := range gf.r.ReadAll(ctx) {

					if _, ok := todayRecipeIDs[recipe.ID]; ok {
						continue
					}

					if recipe.Calories+dayCalories > MaxCalories {
						continue
					}

					week.Menu[i].Recipes[j] = recipe
					id = recipe.ID
					break
				}
			} else {
				week.Menu[i].Recipes[j] = gf.r.Read(ctx, id)
			}
			todayRecipeIDs[id] = week.Menu[i].Recipes[j]
			dayCalories = dayCalories + week.Menu[i].Recipes[j].Calories
		}
	}

	return week
}