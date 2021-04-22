package storage

import (
	"context"

	"github.com/wormi4ok/menuplanner/internal"
)

func (s *DB) preloadCourses() {
	courses := []internal.Course{
		{Name: internal.CourseBreakfast},
		{Name: internal.CourseMain},
		{Name: internal.CoursePudding},
	}
	for _, course := range courses {
		s.db.FirstOrCreate(&course, &course)
	}
}

func (s *DB) ReadCourse(ctx context.Context, id int) (c *internal.Course) {
	s.db.WithContext(ctx).First(&c, id)
	return c
}

func (s *DB) ReadAllCourses(ctx context.Context) (cc []*internal.Course) {
	s.db.WithContext(ctx).Find(&cc)
	return
}
