package mongogo

import "go.mongodb.org/mongo-driver/mongo"

func (m *Mongo[T]) GetCollection() *mongo.Collection {
	return mi.Db.Collection(m.collectionName)
}
