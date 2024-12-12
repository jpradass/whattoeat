package models

type Ingredient struct {
	Name string
}

type Step struct {
	StepDescription string
}

type Recipe struct {
	Id          string
	Title       string
	Image       string
	Description string
	Ingredients []Ingredient
	Steps       []Step
}
