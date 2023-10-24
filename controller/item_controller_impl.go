package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/raafly/inventory-management/helper"
	"github.com/raafly/inventory-management/model"
	portService "github.com/raafly/inventory-management/service/port"
)

type ItemControllerImpl struct {
	ItemService portService.ItemService
}

func NewItemController(itemService portService.ItemService) *ItemControllerImpl {
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
	Id := params.ByName("name")
	
	c.ItemService.Delete(r.Context(), Id)
	webResponse := model.WebResponse {
		Code: 201,
		Status: "SUCCESS",
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c *ItemControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	Id := params.ByName("name")

	item := c.ItemService.FindById(r.Context(), Id)
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