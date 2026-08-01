package mock

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/wormi4ok/menuplanner/internal"
)

type Recipes struct {
	all []*internal.Recipe
}

func (rs *Recipes) ReadRandom(ctx context.Context, course internal.Course, userID int) *internal.Recipe {
	rand.Seed(time.Now().UnixNano())
	var rr []*internal.Recipe
	for _, r := range rs.all {
		for _, c := range r.Courses {
			if course.ID == c.ID {
				rr = append(rr, r)
			}
		}
	}
	if rr == nil {
		return nil
	}
	return rr[rand.Intn(len(rr))]
}

func (rs *Recipes) Create(_ context.Context, r *internal.Recipe) (*internal.Recipe, error) {
	r.ID = len(rs.all) + 1
	rs.all = append(rs.all, r)
	return r, nil
}

func (rs *Recipes) Update(_ context.Context, recipe *internal.Recipe) (*internal.Recipe, error) {
	for i, r := range rs.all {
		if r.ID == recipe.ID {
			rs.all[i] = recipe
			return recipe, nil
		}
	}
	return nil, errors.New("not found")
}

func (rs *Recipes) Read(ctx context.Context, userID int, id int) *internal.Recipe {
	for _, r := range rs.all {
		if r.ID == id {
			return r
		}
	}
	return nil
}

func (rs *Recipes) ReadAll(ctx context.Context, userID int) []*internal.Recipe {
	return rs.all
}

func (rs *Recipes) Delete(_ context.Context, id int) bool {
	for i, r := range rs.all {
		if r.ID == id {
			rs.all[i] = rs.all[len(rs.all)-1]
			rs.all = rs.all[:len(rs.all)-1]
			return true
		}
	}
	return false
}

type Weeks struct {
	Recipes internal.RecipeRepository
	current *internal.Week
}

func (ws *Weeks) UpdateCurrent(ctx context.Context, userID int, week *internal.Week) *internal.Week {
	ws.current = week
	return ws.ReadCurrent(ctx, userID)
}

func (ws *Weeks) ReadCurrent(ctx context.Context, userID int) *internal.Week {
	if ws.current == nil {
		ws.current = &internal.Week{Menu: map[int]*internal.DailyMenu{}}
	}
	c := ws.current
	for i, day := range c.Menu {
		for k, recipe := range day.Recipes {
			if recipe.IsEmpty() {
				continue
			}
			if full := ws.Recipes.Read(ctx, userID, recipe.ID); full != nil {
				c.Menu[i].Recipes[k] = *full
			} else {
				delete(c.Menu[i].Recipes, k)
			}
		}
	}

	return c
}

func (ws *Weeks) DeleteSlot(_ context.Context, _ int, _ int, day int, slot int) error {
	if ws.current == nil {
		return nil
	}
	if menu, exists := ws.current.Menu[day]; exists {
		delete(menu.Recipes, slot)
	}
	return nil
}

type Users struct {
	all []*internal.User
}

func (us *Users) CreateUser(_ context.Context, user *internal.User) error {
	user.ID = len(us.all) + 1
	us.all = append(us.all, user)
	return nil
}

func (us *Users) UpdateUser(_ context.Context, user *internal.User) error {
	for i, u := range us.all {
		if u.ID == user.ID {
			us.all[i] = user
			return nil
		}
	}
	return internal.NewError(errors.New("user not found"), internal.ErrorNotFound)
}

func (us *Users) ReadUser(_ context.Context, id int) (*internal.User, error) {
	for _, u := range us.all {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, internal.NewError(errors.New("user not found"), internal.ErrorNotFound)
}

func (us *Users) ReadUserByEmail(_ context.Context, email string) (*internal.User, error) {
	for _, u := range us.all {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, internal.NewError(errors.New("user not found"), internal.ErrorNotFound)
}

type Courses struct {
	all []*internal.Course
}

func NewCourses() *Courses {
	return &Courses{all: []*internal.Course{
		{ID: 1, Name: internal.CourseBreakfast},
		{ID: 2, Name: internal.CourseMain},
		{ID: 3, Name: internal.CoursePudding},
	}}
}

func (cs *Courses) ReadCourse(_ context.Context, id int) *internal.Course {
	for _, c := range cs.all {
		if c.ID == id {
			return c
		}
	}
	return nil
}

func (cs *Courses) ReadAllCourses(_ context.Context) []*internal.Course {
	return cs.all
}
