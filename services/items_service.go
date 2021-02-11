package services

import (
	"github.com/Seanlinsanity/GolangElasticSearchAPI/domain/items"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, errors.ApiError)
	Get(string) (*items.Item, errors.ApiError)
}

type itemsService struct {
}

func (service *itemsService) Create(item items.Item) (*items.Item, errors.ApiError) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (service *itemsService) Get(id string) (*items.Item, errors.ApiError) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}
