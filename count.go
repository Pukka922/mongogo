package mongogo

import "context"

func (m *Mongo[T]) Count(query interface{}) (int64, error) {
	collection := m.GetCollection()

	count, err := collection.CountDocuments(context.Background(), query)

	if err != nil {
		return 0, err
	}

	return count, nil
}
