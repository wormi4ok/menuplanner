package mock

import (
	"context"

	"github.com/wormi4ok/menuplanner/internal"
)

type Recipes struct {
	all []*internal.Recipe
}

func (rs *Recipes) Create(_ context.Context, r *internal.Recipe) (*internal.Recipe, error) {
	r.ID = len(rs.all) + 1
	rs.all = append(rs.all, r)
	return r, nil
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

func (ws *Weeks) UpdateCurrent(week *internal.Week) *internal.Week {
	ws.current = week
	return ws.current
}

func (ws *Weeks) ReadCurrent() *internal.Week {
	c := ws.current
	for i, day := range c.Menu {
		for k, recipe := range day.Recipes {
			r := ws.Recipes.Read(context.TODO(), recipe.ID)
			c.Menu[i].Recipes[k] = r
		}
	}

	return c
}
