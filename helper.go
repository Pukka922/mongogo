package mongogo

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func createSelectFilter(filtersRaw string) bson.D {
	filters := bson.D{}
	split := strings.Split(filtersRaw, " ")

	for _, value := range split {
		if strings.HasPrefix(value, "-") {
			filters = append(filters, bson.E{Key: strings.Trim(value, "-"), Value: 0})
		} else {
			filters = append(filters, bson.E{Key: value, Value: 1})
		}
	}

	return filters
}
