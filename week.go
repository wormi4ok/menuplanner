package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wormi4ok/menuplanner/internal"
)

func weekHandler() http.HandlerFunc {
	r1 := internal.Recipe{
		ID:       1,
		Name:     "Moroccan Carrot Soup",
		Calories: 300,
		Protein:  60,
		Fat:      50,
		Carbs:    50,
	}

	r2 := internal.Recipe{
		ID:       2,
		Name:     "Ovsyanoblin",
		Calories: 350,
		Protein:  50,
		Fat:      10,
		Carbs:    30,
	}

	r3 := internal.Recipe{
		ID:       3,
		Name:     "Pasta Carbonara",
		Calories: 500,
		Protein:  120,
		Fat:      80,
		Carbs:    260,
	}

	d1 := map[int]internal.Recipe{0: r1, 1: r2}
	d2 := map[int]internal.Recipe{2: r3}
	m := internal.Menu{0: internal.DailyMenu{Recipes: d1}, 1: internal.DailyMenu{d2}}
	week := internal.Week{Menu: m}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(week)
		if err != nil {
			log.Printf("Handler error: %v", err)
			w.WriteHeader(500)
		}
	}
}
