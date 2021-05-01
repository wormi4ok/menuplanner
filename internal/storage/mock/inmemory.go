package mock

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/wormi4ok/menuplanner/internal"
)

type Recipes struct {
	all []*internal.Recipe
}

func (rs *Recipes) ReadRandom(_ context.Context, course internal.Course) *internal.Recipe {
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

func (rs *Recipes) LoadFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var rr []*internal.Recipe

	if err = json.Unmarshal(data, &rr); err != nil {
		return err
	}
	rs.all = rr

	return nil
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

func (rs *Recipes) Read(_ context.Context, id int) *internal.Recipe {
	for _, r := range rs.all {
		if r.ID == id {
			return r
		}
	}
	return nil
}

func (rs *Recipes) ReadAll(_ context.Context) []*internal.Recipe {
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

func (ws *Weeks) LoadFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var w *internal.Week

	if err = json.Unmarshal(data, &w); err != nil {
		return err
	}
	ws.current = w

	return nil
}

func (ws *Weeks) UpdateCurrent(_ context.Context, week *internal.Week) *internal.Week {
	ws.current = week
	return ws.current
}

func (ws *Weeks) ReadCurrent(_ context.Context) *internal.Week {
	c := ws.current
	for i, day := range c.Menu {
		for k, recipe := range day.Recipes {
			c.Menu[i].Recipes[k] = *ws.Recipes.Read(context.TODO(), recipe.ID)
		}
	}

	return c
}

func (ws *Weeks) DeleteSlot(_ context.Context, _, day, slot int) error {
	menu := ws.current.Menu
	if menu != nil {
		if _, exists := menu[day]; exists {
			if _, exists = menu[day].Recipes[slot]; exists {
				ws.current.Menu[day].Recipes[slot] = internal.Recipe{}
				return nil
			}
		}
	}
	return errors.New("not found")
}
