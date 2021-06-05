package storage

import (
	"context"

	"github.com/wormi4ok/menuplanner/internal"
	"gorm.io/gorm/clause"
)

func (s *DB) Create(ctx context.Context, recipe *internal.Recipe) (*internal.Recipe, error) {
	return recipe, s.db.WithContext(ctx).Create(recipe).Error
}

func (s *DB) Update(ctx context.Context, recipe *internal.Recipe) (*internal.Recipe, error) {
	oldRecipe := internal.Recipe{ID: recipe.ID}
	if err := s.db.WithContext(ctx).Model(&oldRecipe).Association("Courses").Clear(); err != nil {
		return nil, err
	}

	if res := s.db.WithContext(ctx).Save(recipe); res.Error != nil {
		return nil, res.Error
	}

	return recipe, nil
}

func (s *DB) Delete(ctx context.Context, id int) bool {
	r := &internal.Recipe{ID: id}
	return s.db.WithContext(ctx).Model(&r).Select(clause.Associations).Delete(&r).Error == nil
}

func (s *DB) Read(ctx context.Context, userID int, id int) *internal.Recipe {
	r := &internal.Recipe{}
	s.db.WithContext(ctx).Preload("Courses").First(&r, "user_id = ? AND id = ?", userID, id)
	return r
}

func (s *DB) ReadAll(ctx context.Context, userID int) (rr []*internal.Recipe) {
	s.db.WithContext(ctx).Preload("Courses").Find(&rr, "user_id = ?", userID)
	return
}

func (s *DB) ReadRandom(ctx context.Context, course internal.Course, userID int) (r *internal.Recipe) {
	var id int
	s.db.
		Table("recipe_courses").
		Select("recipe_id").
		Where("course_id = ?", course.ID).
		Joins("JOIN recipes ON recipes.id = recipe_courses.recipe_id AND recipes.user_id = ?", userID).
		Order("RAND()").
		Scan(&id)
	s.db.WithContext(ctx).Preload("Courses").First(&r, id)
	return r
}
