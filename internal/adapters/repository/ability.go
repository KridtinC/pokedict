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

type abilityRepository struct {
	db             *mongo.Collection
	collectionName string
}

func NewAbilityRepository(db *mongo.Database) ports.AbilityRepository {
	return &abilityRepository{
		db:             db.Collection(collectionNameAbilities),
		collectionName: collectionNameAbilities,
	}
}

func (r *abilityRepository) Find(ctx context.Context, page int, limit int, offset int) ([]domain.Ability, error) {
	panic("implement me")
}

func (r *abilityRepository) FindAll(ctx context.Context) ([]domain.Ability, error) {
	panic("implement me")
}

func (r *abilityRepository) FindByID(ctx context.Context, id int) (domain.Ability, error) {
	var ability model.Ability
	result := r.db.FindOne(ctx, bson.M{"id": id})
	if err := result.Decode(&ability); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Ability{}, errmsg.ErrNotFound
		}
		return domain.Ability{}, err
	}
	return ability.ToDomain(), nil
}

func (r *abilityRepository) FindByIDs(ctx context.Context, ids []int) ([]domain.Ability, error) {
	var abilities []domain.Ability
	cursor, err := r.db.Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		return abilities, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var ability model.Ability
		if err := cursor.Decode(&ability); err != nil {
			return abilities, err
		}
		abilities = append(abilities, ability.ToDomain())
	}
	return abilities, nil
}

func (r *abilityRepository) Insert(ctx context.Context, ability domain.Ability) error {
	data := model.NewAbilityFromDomain(ability)
	_, err := r.db.InsertOne(ctx, data)
	return err
}
