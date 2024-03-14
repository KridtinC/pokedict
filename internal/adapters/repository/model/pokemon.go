package model

import "github.com/KridtinC/pokedict/internal/core/domain"

type Pokemon struct {
	ID        int    `bson:"id"`
	Name      string `bson:"name"`
	BaseExp   int    `bson:"base_experience"`
	Height    int    `bson:"height"`
	Weight    int    `bson:"weight"`
	Type1     string `bson:"type1"`
	Type2     string `bson:"type2"`
	HP        int    `bson:"hp"`
	Attack    int    `bson:"attack"`
	Defense   int    `bson:"defense"`
	SpAttack  int    `bson:"sp_attack"`
	SpDefense int    `bson:"sp_defense"`
	Speed     int    `bson:"speed"`
	Abilities []int  `bson:"abilities"`
	Moves     []int  `bson:"moves"`
}

func (p *Pokemon) ToDomain() domain.Pokemon {
	var pkm = domain.Pokemon{
		ID:        p.ID,
		Name:      p.Name,
		BaseExp:   p.BaseExp,
		Height:    p.Height,
		Weight:    p.Weight,
		Type1:     domain.Type(p.Type1),
		Type2:     domain.Type(p.Type2),
		HP:        p.HP,
		Attack:    p.Attack,
		Defense:   p.Defense,
		SpAttack:  p.SpAttack,
		SpDefense: p.SpDefense,
		Speed:     p.Speed,
	}
	for _, a := range p.Abilities {
		pkm.Abilities = append(pkm.Abilities, domain.Ability{ID: a})
	}
	for _, m := range p.Moves {
		pkm.Moves = append(pkm.Moves, domain.Move{ID: m})
	}
	return pkm
}

func NewPokemonFromDomain(p domain.Pokemon) Pokemon {
	var pkm = Pokemon{
		ID:        p.ID,
		Name:      p.Name,
		BaseExp:   p.BaseExp,
		Height:    p.Height,
		Weight:    p.Weight,
		Type1:     string(p.Type1),
		Type2:     string(p.Type2),
		HP:        p.HP,
		Attack:    p.Attack,
		Defense:   p.Defense,
		SpAttack:  p.SpAttack,
		SpDefense: p.SpDefense,
		Speed:     p.Speed,
	}
	for _, a := range p.Abilities {
		pkm.Abilities = append(pkm.Abilities, a.ID)
	}
	for _, m := range p.Moves {
		pkm.Moves = append(pkm.Moves, m.ID)
	}
	return pkm
}
