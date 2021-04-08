package main

type Recipe struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"imageUrl"`

	Calories int `json:"calories"`
	Protein  int `json:"protein"`
	Fat      int `json:"fat"`
	Carbs    int `json:"carbs"`
}

type DailyMenu struct {
	Recipes map[int]Recipe `json:"recipes"`
}

type Menu map[int]DailyMenu

type Week struct {
	Menu Menu `json:"menu"`
}
