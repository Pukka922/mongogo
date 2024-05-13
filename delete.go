package mongogo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *Mongo[T]) Delete(query interface{}) (bool, error) {
	collection := m.GetCollection()

	_, err := collection.DeleteOne(context.Background(), query)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *Mongo[T]) DeleteById(id string) (bool, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return false, err
	}

	return m.Delete(bson.D{
		{Key: "_id", Value: objectId},
	})
}
