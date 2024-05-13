package mongogo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *Mongo[T]) Update(query, data interface{}, filter string) (*T, error) {
	collection := m.GetCollection()

	_, err := collection.UpdateOne(context.Background(), query, data)

	if err != nil {
		return nil, err
	}

	return m.FindOne(query, filter)
}

func (m *Mongo[T]) UpdateOneWithReplace(query, data interface{}, filter string) (*T, error) {
	collection := m.GetCollection()

	_, err := collection.ReplaceOne(context.Background(), query, data)

	if err != nil {
		return nil, err
	}

	return m.FindOne(query, filter)
}

func (m *Mongo[T]) UpdateOneWithReplaceById(id string, data interface{}, filter string) (*T, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	return m.UpdateOneWithReplace(bson.D{{Key: "_id", Value: objectId}}, data, filter)
}
