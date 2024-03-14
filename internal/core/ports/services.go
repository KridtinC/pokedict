package ports

import (
	"context"

	"github.com/KridtinC/pokedict/internal/core/domain"
)

type PokemonService interface {
	FindAll(ctx context.Context) ([]domain.Pokemon, error)
	FindByID(ctx context.Context, id int) (domain.Pokemon, error)
}
