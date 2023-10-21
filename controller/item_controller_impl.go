package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/raafly/inventory-management/helper"
	"github.com/raafly/inventory-management/model"
	"github.com/raafly/inventory-management/service"
)

type ItemControllerImpl struct {
	ItemService service.ItemServiceImpl
}

func NewItemController(itemService service.ItemServiceImpl) *ItemControllerImpl {
	return &ItemControllerImpl{
		ItemService: itemService,
	}
}

func (c *ItemControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	itemCreateRequest := model.ItemCreate{}
	helper.ReadFromRequestBody(r, &itemCreateRequest)

	item := c.ItemService.Create(r.Context(), itemCreateRequest)
	webResponse := model.WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: item,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c *ItemControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	itemCreateRequest := model.ItemUpdate{}
	helper.ReadFromRequestBody(r, &itemCreateRequest)

	item := c.ItemService.Update(r.Context(), itemCreateRequest)
	webResponse := model.WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: item,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c *ItemControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := r.URL.Query().Get("userId")
	userId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	c.ItemService.Delete(r.Context(), userId)
	webResponse := model.WebResponse {
		Code: 201,
		Status: "SUCCESS",
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c *ItemControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := r.URL.Query().Get("userId")
	userId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	item := c.ItemService.FindById(r.Context(), userId)
	webResponse := model.WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: item,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c *ItemControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	items := c.ItemService.FindAll(r.Context())
	webResponse := model.WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: items,
	}

	helper.WriteToRequestBody(w, webResponse)
}