package ports

import (
	"context"

	"github.com/KridtinC/pokedict/internal/core/domain"
)

type TxRepository interface {
	ProcessTx(ctx context.Context, fn func(ctx context.Context) error) (any, error)
}

type PokemonRepository interface {
	Find(ctx context.Context, page int, limit int, offset int) ([]domain.Pokemon, error)
	FindAll(ctx context.Context) ([]domain.Pokemon, error)
	FindByID(ctx context.Context, id int) (domain.Pokemon, error)
	Insert(ctx context.Context, pkm domain.Pokemon) error
}

type MoveRepository interface {
	Find(ctx context.Context, page int, limit int, offset int) ([]domain.Move, error)
	FindAll(ctx context.Context) ([]domain.Move, error)
	FindByID(ctx context.Context, id int) (domain.Move, error)
	FindByIDs(ctx context.Context, ids []int) ([]domain.Move, error)
	Insert(ctx context.Context, move domain.Move) error
}

type AbilityRepository interface {
	Find(ctx context.Context, page int, limit int, offset int) ([]domain.Ability, error)
	FindAll(ctx context.Context) ([]domain.Ability, error)
	FindByID(ctx context.Context, id int) (domain.Ability, error)
	FindByIDs(ctx context.Context, ids []int) ([]domain.Ability, error)
	Insert(ctx context.Context, ability domain.Ability) error
}
