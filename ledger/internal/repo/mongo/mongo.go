package mongo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"ledger/internal/helper/di"
	"ledger/internal/repo"

	"github.com/samber/do"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepo struct {
	col *mongo.Collection
	ctx context.Context
}

func NewLedgerRepo(i *do.Injector) (repo.LedgerRepo, error) {
	var cfg MongoConfig
	if err := di.GetConfigFromDI(i, &cfg); err != nil {
		return nil, errors.Join(errGetCfg, err)
	}
	mongoURI := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d", cfg.Username, cfg.Password, cfg.Host, cfg.Port)

	// Set client options and connect
	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx := context.Background()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the MongoDB server to verify connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}
	fmt.Println("Successfully connected to MongoDB")

	return mongoRepo{
		col: client.Database("ledger").Collection("transactions"),
		ctx: ctx,
	}, nil
}

func (mr mongoRepo) WriteLogs(data []byte) error {
	var transaction transaction
	err := json.Unmarshal(data, &transaction)
	if err != nil {
		return fmt.Errorf("failed to unmarshal transaction: %w", err)
	}
	_, err = mr.col.InsertOne(mr.ctx, transaction)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	return nil
}
