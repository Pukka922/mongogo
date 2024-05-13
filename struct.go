package mongogo

type Mongo[T any] struct {
	collectionName string
}

func NewMongo[T any](collectionName string) *Mongo[T] {
	return &Mongo[T]{
		collectionName: collectionName,
	}
}
