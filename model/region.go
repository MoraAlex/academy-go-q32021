package model

type Region struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Generation Generation `json:"main_generation"`
}
