package internal

type WeekRepository interface {
	ReadCurrent() *Week
	UpdateCurrent(*Week) *Week
}

type DailyMenu struct {
	Recipes map[int]Recipe `json:"recipes"`
}

type Menu map[int]DailyMenu

type Week struct {
	Menu Menu `json:"menu"`
}
