package items

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Seanlinsanity/GolangElasticSearchAPI/clients/elasticSearch"
	"github.com/Seanlinsanity/GolangElasticSearchAPI/domain/queries"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() errors.ApiError {
	result, err := elasticSearch.Client.Index(indexItems, i)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("save es client db err %s", err.Error()))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() errors.ApiError {
	itemId := i.Id
	result, err := elasticSearch.Client.Get(indexItems, itemId)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return errors.NewNotFoundError(fmt.Sprintf("no item founc with id: %s", i.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("get es client id err: %s", err.Error()))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return errors.NewInternalServerError("error when trying to marshal es response")
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return errors.NewInternalServerError("error when trying to parse es response")
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, errors.ApiError) {
	result, err := elasticSearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("failed to search documents err: %s", err.Error()))
	}
	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, errors.NewInternalServerError(fmt.Sprintf("failed to parse item err: %s", err.Error()))
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, errors.NewNotFoundError("not items found for giving criteria")
	}

	return items, nil
}
