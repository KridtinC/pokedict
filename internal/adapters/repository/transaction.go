package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type txRepository struct {
	db *mongo.Client
}

func NewTxRepository(db *mongo.Client) *txRepository {
	return &txRepository{
		db: db,
	}
}

func (t *txRepository) ProcessTx(ctx context.Context, fn func(ctx context.Context) error) (any, error) {
	ss, err := t.db.StartSession()
	if err != nil {
		return nil, err
	}
	defer ss.EndSession(ctx)
	result, err := ss.WithTransaction(ctx, func(ssctx mongo.SessionContext) (interface{}, error) {
		return nil, fn(ssctx)
	})
	return result, err
}
