package mongogo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Mongo[T]) Aggregate(pipeline mongo.Pipeline, filter string) ([]T, error) {
	ctx := context.Background()
	collection := m.GetCollection()

	if len(filter) > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$project", Value: createSelectFilter(filter)}})
	}

	cursor, err := collection.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var results []T

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
