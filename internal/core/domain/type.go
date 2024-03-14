package domain

type Type string

func (t Type) String() string {
	return string(t)
}

const (
	None     Type = "none"
	Normal   Type = "normal"
	Fighting Type = "fighting"
	Flying   Type = "flying"
	Poison   Type = "poison"
	Ground   Type = "ground"
	Rock     Type = "rock"
	Bug      Type = "bug"
	Ghost    Type = "ghost"
	Steel    Type = "steel"
	Fire     Type = "fire"
	Water    Type = "water"
	Grass    Type = "grass"
	Electric Type = "electric"
	Psychic  Type = "psychic"
	Ice      Type = "ice"
	Dragon   Type = "dragon"
	Dark     Type = "dark"
	Fairy    Type = "fairy"
)
