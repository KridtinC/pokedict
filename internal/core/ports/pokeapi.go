package ports

import (
	"context"

	"github.com/KridtinC/pokedict/internal/core/domain"
)

type PokeAPI interface {
	GetPokemonByID(ctx context.Context, id int) (domain.Pokemon, error)
}
