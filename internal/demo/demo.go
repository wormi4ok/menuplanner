package demo

import (
	"context"
	_ "embed"
	"encoding/json"

	"github.com/wormi4ok/menuplanner/internal"
)

const (
	name     = "Demo User"
	email    = "demo@demo.com"
	password = "demo"
)

//go:embed recipes.json
var demoRecipes []byte

//go:embed week.json
var demoWeek []byte

func PreloadData(
	ctx context.Context,
	us internal.UserRepository,
	rs internal.RecipeRepository,
	ws internal.WeekRepository,
) error {
	user, err := writeUser(ctx, us)
	if err != nil {
		return err
	}

	if err := writeRecipes(ctx, rs, user.ID); err != nil {
		return err
	}

	if err := writeWeek(ctx, ws, user.ID); err != nil {
		return err
	}

	return nil
}

func writeUser(ctx context.Context, us internal.UserRepository) (user *internal.User, err error) {
	if user, _ = us.ReadUserByEmail(ctx, email); user == nil {
		if user, err = internal.NewUser(email, password); err != nil {
			return nil, err
		}
		user.Name = name
		err = us.CreateUser(ctx, user)
	}

	return user, err
}

func writeRecipes(ctx context.Context, rs internal.RecipeRepository, userID int) error {
	existingRecipes := rs.ReadAll(ctx, userID)

	var rr []*internal.Recipe
	if err := json.Unmarshal(demoRecipes, &rr); err != nil {
		return err
	}

	for _, recipe := range rr {
		recipe.UserID = userID
		found := false
		for _, er := range existingRecipes {
			if er.ID == recipe.ID {
				found = true
				if _, err := rs.Update(ctx, recipe); err != nil {
					return err
				}
				break
			}
		}

		if !found {
			if _, err := rs.Create(ctx, recipe); err != nil {
				return err
			}
		}
	}

	for _, recipe := range existingRecipes {
		found := false
		for _, r := range rr {
			if recipe.ID == r.ID {
				found = true
				break
			}
		}
		if !found {
			rs.Delete(ctx, recipe.ID)
		}
	}

	return nil
}

func writeWeek(ctx context.Context, ws internal.WeekRepository, userID int) error {
	var w *internal.Week
	if err := json.Unmarshal(demoWeek, &w); err != nil {
		return err
	}

	// set default data
	ws.UpdateCurrent(ctx, userID, w)

	// remove everything else
	_ = ws.DeleteSlot(ctx,userID,1,2,0)
	for i := 3; i < 7; i++ {
		for j := 0; j < 3; j++ {
			if err := ws.DeleteSlot(ctx, userID, 1, i, j); err != nil {
				return err
			}
		}
	}
	return nil
}
