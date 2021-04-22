package internal

import "context"

const (
	CourseBreakfast = "breakfast"
	CourseMain      = "main"
	CoursePudding   = "pudding"
)

type CourseReader interface {
	ReadCourse(ctx context.Context, id int) *Course
	ReadAllCourses(ctx context.Context) []*Course
}

type Course struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"not null;size:255"`
}
