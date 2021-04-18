package storage

import (
	"context"

	"github.com/wormi4ok/menuplanner/internal"
)

func (s *DB) Create(ctx context.Context, recipe *internal.Recipe) (*internal.Recipe, error) {
	return recipe, s.db.WithContext(ctx).Create(recipe).Error
}

func (s *DB) Delete(ctx context.Context, id int) bool {
	return s.db.WithContext(ctx).Delete(&internal.Recipe{}, id).Error == nil
}

func (s *DB) Read(ctx context.Context, id int) *internal.Recipe {
	r := &internal.Recipe{}
	s.db.WithContext(ctx).First(&r, id)
	return r
}

func (s *DB) ReadAll(ctx context.Context) (rr []*internal.Recipe) {
	s.db.WithContext(ctx).Find(&rr)
	return
}
