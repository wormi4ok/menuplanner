package storage

import (
	"context"
	"log"

	"github.com/wormi4ok/menuplanner/internal"
	"gorm.io/gorm/clause"
)

const currentWeek = 1

type Week struct {
	ID       int `gorm:"primaryKey;autoIncrement:false"`
	Day      int `gorm:"primaryKey;autoIncrement:false"`
	Slot     int `gorm:"primaryKey;autoIncrement:false"`
	RecipeID int
	Recipe   internal.Recipe
}

func (s *DB) ReadCurrent(ctx context.Context) *internal.Week {
	var week internal.Week
	week.Menu = make(map[int]internal.DailyMenu, 7)

	var w []Week
	s.db.WithContext(ctx).Joins("Recipe").Find(&w, "weeks.id = ?", currentWeek)
	for _, line := range w {
		if week.Menu[line.Day].Recipes == nil {
			week.Menu[line.Day] = internal.DailyMenu{Recipes: map[int]*internal.Recipe{
				line.Slot: &line.Recipe,
			}}
		} else {
			week.Menu[line.Day].Recipes[line.Slot] = &line.Recipe
		}
	}

	return &week
}

func (s *DB) UpdateCurrent(ctx context.Context, week *internal.Week) *internal.Week {
	for day, menu := range week.Menu {
		for slot, recipe := range menu.Recipes {
			if recipe == nil || recipe.ID == 0 {
				continue
			}
			w := &Week{
				ID:       currentWeek,
				Day:      day,
				Slot:     slot,
				RecipeID: recipe.ID,
			}
			result := s.db.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(w)

			if result.Error != nil {
				log.Printf("Error wle insering: %s", result.Error)
			}
		}
	}
	return s.ReadCurrent(nil)
}

func (s *DB) DeleteSlot(ctx context.Context, week, day, slot int) error {
	w := &Week{
		ID:   week,
		Day:  day,
		Slot: slot,
	}
	return s.db.WithContext(ctx).Delete(&w).Error
}
