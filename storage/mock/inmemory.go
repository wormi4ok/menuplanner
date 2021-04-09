package mock

import (
	"context"

	"github.com/wormi4ok/menuplanner/internal"
)

type Recipes struct {
	all []*internal.Recipe
}

func (rs *Recipes) Create(_ context.Context, r *internal.Recipe) *internal.Recipe {
	r.ID = len(rs.all) + 1
	rs.all = append(rs.all, r)
	return r
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
