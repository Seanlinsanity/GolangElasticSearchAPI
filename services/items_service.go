package services

import (
	"net/http"

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

func (service *itemsService) Create(items.Item) (*items.Item, errors.ApiError) {
	return nil, errors.NewApiError(http.StatusNotImplemented, "need implement!")
}

func (service *itemsService) Get(string) (*items.Item, errors.ApiError) {
	return nil, errors.NewApiError(http.StatusNotImplemented, "need implement!")
}
