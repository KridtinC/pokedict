package model

import "github.com/KridtinC/pokedict/internal/core/domain"

type FindByIDResponse struct {
	Data Pokemon `json:"data"`
}

type Pokemon struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BaseExp   int       `json:"base_exp"`
	Height    int       `json:"height"`
	Weight    int       `json:"weight"`
	Type1     string    `json:"type1"`
	Type2     string    `json:"type2"`
	HP        int       `json:"hp"`
	Attack    int       `json:"attack"`
	Defense   int       `json:"defense"`
	SpAttack  int       `json:"sp_attack"`
	SpDefense int       `json:"sp_defense"`
	Speed     int       `json:"speed"`
	Abilities []Ability `json:"abilities"`
}

func PokemonFromDomain(pkm domain.Pokemon) Pokemon {
	return Pokemon{
		ID:        pkm.ID,
		Name:      pkm.Name,
		BaseExp:   pkm.BaseExp,
		Height:    pkm.Height,
		Weight:    pkm.Weight,
		Type1:     pkm.Type1.String(),
		Type2:     pkm.Type2.String(),
		HP:        pkm.HP,
		Attack:    pkm.Attack,
		Defense:   pkm.Defense,
		SpAttack:  pkm.SpAttack,
		SpDefense: pkm.SpDefense,
		Speed:     pkm.Speed,
		Abilities: AbilitiesFromDomain(pkm.Abilities),
	}
}

func AbilitiesFromDomain(abilities []domain.Ability) []Ability {
	var result []Ability
	for _, ability := range abilities {
		result = append(result, Ability{
			ID:       ability.ID,
			Name:     ability.Name,
			IsHidden: ability.IsHidden,
		})
	}
	return result
}
