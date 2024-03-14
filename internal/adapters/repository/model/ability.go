package model

import "github.com/KridtinC/pokedict/internal/core/domain"

type Ability struct {
	ID       int    `bson:"id"`
	Name     string `bson:"name"`
	IsHidden bool   `bson:"is_hidden"`
}

func (a *Ability) ToDomain() domain.Ability {
	return domain.Ability{
		ID:       a.ID,
		Name:     a.Name,
		IsHidden: a.IsHidden,
	}
}

func NewAbilityFromDomain(a domain.Ability) Ability {
	return Ability{
		ID:       a.ID,
		Name:     a.Name,
		IsHidden: a.IsHidden,
	}
}
