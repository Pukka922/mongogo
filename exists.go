package mongogo

func (m *Mongo[T]) Exists(query interface{}) (bool, error) {
	count, err := m.Count(query)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
