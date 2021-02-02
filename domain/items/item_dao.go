package items

import (
	"github.com/Seanlinsanity/GolangElasticSearchAPI/clients/elasticSearch"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() errors.ApiError {
	result, err := elasticSearch.Client.Index(indexItems, i)
	if err != nil {
		return errors.NewInternalServerError("error when trying to save item, database error")
	}
	i.Id = result.Id
	return nil
}
