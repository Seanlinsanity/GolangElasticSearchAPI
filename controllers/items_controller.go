package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Seanlinsanity/GolangElasticSearchAPI/domain/items"
	"github.com/Seanlinsanity/GolangElasticSearchAPI/domain/queries"
	"github.com/Seanlinsanity/GolangElasticSearchAPI/services"
	"github.com/Seanlinsanity/GolangElasticSearchAPI/utils/http_utils"
	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	reqSellerID := r.Header.Get("Seller-Id")
	sellerID, parseErrr := strconv.ParseInt(reqSellerID, 10, 64)
	if parseErrr != nil {
		http_utils.RespondErr(w, errors.NewBadRequestError("invalid seller id"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := errors.NewBadRequestError("invalid request body")
		http_utils.RespondErr(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := errors.NewBadRequestError("invalid json body")
		http_utils.RespondErr(w, respErr)
		return
	}

	itemRequest.Seller = sellerID

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondErr(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemsId := strings.TrimSpace((vars["id"]))
	item, err := services.ItemsService.Get(itemsId)
	if err != nil {
		http_utils.RespondErr(w, err)
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		http_utils.RespondErr(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		http_utils.RespondErr(w, apiErr)
		return
	}

	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.RespondErr(w, searchErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, items)
}
