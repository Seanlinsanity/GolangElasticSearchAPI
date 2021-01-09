package controllers

import (
	"net/http"
	"strconv"

	"github.com/Seanlinsanity/GolangElasticSearchAPI/domain/items"
	"github.com/Seanlinsanity/GolangElasticSearchAPI/services"
	"github.com/Seanlinsanity/GolangElasticSearchAPI/utils/http_utils"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Header.Get("Seller-Id")
	i, parseErrr := strconv.ParseInt(sellerID, 10, 64)
	if parseErrr != nil {
		http_utils.RespondErr(w, errors.NewBadRequestError("invalid seller id"))
		return
	}

	item := items.Item{
		Seller: i,
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		http_utils.RespondErr(w, err)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	_, err := services.ItemsService.Get("123")
	if err != nil {
		http_utils.RespondErr(w, err)
	}
}
