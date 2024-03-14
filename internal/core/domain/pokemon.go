package domain

type Pokemon struct {
	ID        int
	Name      string
	BaseExp   int
	Height    int
	Weight    int
	Type1     Type
	Type2     Type
	HP        int
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
	Abilities []Ability
	Moves     []Move
}
