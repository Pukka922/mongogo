package mongogo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *Mongo[T]) InsertOne(newDoc T, filter string) (*T, error) {
	collection := m.GetCollection()

	result, err := collection.InsertOne(context.Background(), newDoc)

	if err != nil {
		return nil, err
	}

	return m.FindOne(bson.D{
		{Key: "_id", Value: result.InsertedID},
	}, "")
}
