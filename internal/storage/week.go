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
	UserID   int `gorm:"primaryKey;autoIncrement:false;default:1"`
	User     internal.User
	RecipeID int
	Recipe   internal.Recipe
}

func (s *DB) ReadCurrent(ctx context.Context, userID int) *internal.Week {
	var week internal.Week
	week.Menu = make(map[int]*internal.DailyMenu, 7)

	var w []Week
	s.db.WithContext(ctx).Joins("Recipe").Find(&w, "weeks.id = ? AND weeks.user_id = ?", currentWeek, userID)
	for _, line := range w {
		if week.Menu[line.Day] == nil {
			week.Menu[line.Day] = &internal.DailyMenu{Recipes: map[int]internal.Recipe{
				line.Slot: line.Recipe,
			}}
		} else {
			week.Menu[line.Day].Recipes[line.Slot] = line.Recipe
		}
	}

	return &week
}

func (s *DB) UpdateCurrent(ctx context.Context, userID int, week *internal.Week) *internal.Week {
	for day, menu := range week.Menu {
		for slot, recipe := range menu.Recipes {
			if recipe.IsEmpty() {
				continue
			}
			w := &Week{
				ID:       currentWeek,
				Day:      day,
				Slot:     slot,
				UserID:   userID,
				RecipeID: recipe.ID,
			}
			result := s.db.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(w)

			if result.Error != nil {
				log.Printf("Error while insering: %s", result.Error)
			}
		}
	}
	return s.ReadCurrent(ctx, userID)
}

func (s *DB) DeleteSlot(ctx context.Context, userID int, week int, day int, slot int) error {
	w := &Week{
		ID:     week,
		Day:    day,
		Slot:   slot,
		UserID: userID,
	}
	return s.db.WithContext(ctx).Delete(&w).Error
}
