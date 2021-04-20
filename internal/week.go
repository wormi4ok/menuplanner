package internal

import "context"

type (
	WeekRepository interface {
		ReadCurrent(context.Context) *Week
		UpdateCurrent(context.Context, *Week) *Week
		DeleteSlot(ctx context.Context, week, day, slot int) error
	}
)

type DailyMenu struct {
	Recipes map[int]*Recipe `json:"recipes"`
}

type Week struct {
	Menu map[int]DailyMenu `json:"menu"`
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
			recipe, exists := week.Menu[i].Recipes[j]
			if !exists || recipe == nil {
				week.Menu[i].Recipes[j] = &Recipe{}
			}

			id := week.Menu[i].Recipes[j].ID
			if id == 0 {
				for _, recipe := range gf.r.ReadAll(ctx) {

					if _, ok := todayRecipeIDs[recipe.ID]; ok {
						continue
					}

					if recipe.EnergyAmount()+dayCalories > MaxCalories {
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
			dayCalories = dayCalories + week.Menu[i].Recipes[j].EnergyAmount()
		}
	}

	return week
}
