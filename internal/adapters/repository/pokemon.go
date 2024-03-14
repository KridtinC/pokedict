package repository

import (
	"context"

	"github.com/KridtinC/pokedict/internal/adapters/repository/model"
	"github.com/KridtinC/pokedict/internal/core/domain"
	"github.com/KridtinC/pokedict/internal/core/ports"
	"github.com/KridtinC/pokedict/internal/errmsg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type pokemonRepository struct {
	db             *mongo.Collection
	collectionName string
}

func NewPokemonRepository(db *mongo.Database) ports.PokemonRepository {
	return &pokemonRepository{
		db:             db.Collection(collectionNamePokemons),
		collectionName: collectionNamePokemons,
	}
}

func (r *pokemonRepository) Find(ctx context.Context, page int, limit int, offset int) ([]domain.Pokemon, error) {
	panic("implement me")
}

func (r *pokemonRepository) FindAll(ctx context.Context) ([]domain.Pokemon, error) {
	panic("implement me")
}

func (r *pokemonRepository) FindByID(ctx context.Context, id int) (domain.Pokemon, error) {
	var pkm model.Pokemon
	result := r.db.FindOne(ctx, bson.M{"id": id})
	if err := result.Decode(&pkm); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Pokemon{}, errmsg.ErrNotFound
		}
		return domain.Pokemon{}, err
	}
	return pkm.ToDomain(), nil
}

func (r *pokemonRepository) Insert(ctx context.Context, pkm domain.Pokemon) error {
	data := model.NewPokemonFromDomain(pkm)
	_, err := r.db.InsertOne(ctx, data)
	return err
}
