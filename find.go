package mongogo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Mongo[T]) Find(query interface{}, filter string) ([]T, error) {
	collection := m.GetCollection()

	var opts *options.FindOptions

	if len(filter) > 0 {
		opts = options.Find().SetProjection(createSelectFilter(filter))
	}

	cursor, err := collection.Find(context.Background(), query, opts)

	if err != nil {
		return nil, err
	}

	var results []T

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (m *Mongo[T]) FindOne(query interface{}, filter string) (*T, error) {
	collection := m.GetCollection()

	var opts *options.FindOneOptions

	if len(filter) > 0 {
		opts = options.FindOne().SetProjection(createSelectFilter(filter))
	}

	found := collection.FindOne(context.Background(), query, opts)
	entity := new(T)

	err := found.Decode(entity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return entity, nil
}

func (m *Mongo[T]) FindOneById(id, filter string) (*T, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	return m.FindOne(bson.D{
		{Key: "_id", Value: objectId},
	}, "")
}
