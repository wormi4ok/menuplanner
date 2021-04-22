package storage

import (
	"context"

	"github.com/wormi4ok/menuplanner/internal"
	"gorm.io/gorm/clause"
)

func (s *DB) Create(ctx context.Context, recipe *internal.Recipe) (*internal.Recipe, error) {
	return recipe, s.db.WithContext(ctx).Create(recipe).Error
}

func (s *DB) Delete(ctx context.Context, id int) bool {
	r := &internal.Recipe{ID: id}
	return s.db.WithContext(ctx).Model(&r).Select(clause.Associations).Delete(&r).Error == nil
}

func (s *DB) Read(ctx context.Context, id int) *internal.Recipe {
	r := &internal.Recipe{}
	s.db.WithContext(ctx).Preload("Courses").First(&r, id)
	return r
}

func (s *DB) ReadAll(ctx context.Context) (rr []*internal.Recipe) {
	s.db.WithContext(ctx).Preload("Courses").Find(&rr)
	return
}
