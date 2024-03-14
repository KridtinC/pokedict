package services

import (
	"context"

	"github.com/KridtinC/pokedict/internal/core/domain"
	"github.com/KridtinC/pokedict/internal/core/ports"
	"github.com/KridtinC/pokedict/internal/errmsg"
)

type pokemonService struct {
	txRepository      ports.TxRepository
	pokemonRepository ports.PokemonRepository
	abilityRepository ports.AbilityRepository
	pokeAPI           ports.PokeAPI
}

func NewPokemonService(txRepository ports.TxRepository, pokemonRepository ports.PokemonRepository, abilityRepository ports.AbilityRepository, pokeAPI ports.PokeAPI) ports.PokemonService {
	return &pokemonService{
		txRepository,
		pokemonRepository,
		abilityRepository,
		pokeAPI,
	}
}

func (s *pokemonService) FindAll(ctx context.Context) ([]domain.Pokemon, error) {
	return s.pokemonRepository.FindAll(ctx)
}

func (s *pokemonService) FindByID(ctx context.Context, id int) (domain.Pokemon, error) {

	pkm, err := s.pokemonRepository.FindByID(ctx, id)
	if err != nil && err != errmsg.ErrNotFound {
		return domain.Pokemon{}, err
	}

	if err == nil {
		var abiID []int
		for _, abi := range pkm.Abilities {
			abiID = append(abiID, abi.ID)
		}
		abilities, err := s.abilityRepository.FindByIDs(ctx, abiID)
		if err != nil {
			return pkm, err
		}
		pkm.Abilities = abilities
		return pkm, nil
	}

	pkm, err = s.pokeAPI.GetPokemonByID(ctx, id)
	if err != nil {
		return domain.Pokemon{}, err
	}

	if err := s.pokemonRepository.Insert(ctx, pkm); err != nil {
		return pkm, err
	}

	for _, abi := range pkm.Abilities {
		if err := s.abilityRepository.Insert(ctx, abi); err != nil {
			return pkm, err
		}
	}

	return pkm, nil
}
