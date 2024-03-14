package model

type Ability struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsHidden bool   `json:"is_hidden"`
}
