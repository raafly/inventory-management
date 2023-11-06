package listing

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/raafly/inventory-management/helper"
)

type UserController interface{
	SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImpl struct{
	UserService 	UserService
}

func NewUserController(userService UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c *UserControllerImpl) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := UserSignIn{}
	helper.ReadFromRequestBody(r, &userCreateRequest) 

	user, token, err := c.UserService.SignIn(r.Context(), userCreateRequest)	
	helper.PanicIfError(err)
	webResponse := WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: user,
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	helper.WriteToRequestBody(w, webResponse)
}

func (c *UserControllerImpl) SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := UserSignUp{}
	helper.ReadFromRequestBody(r, &userCreateRequest)

	user := c.UserService.SignUp(r.Context(), userCreateRequest)
	webResponse := WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: user,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := params.ByName("username")
	c.UserService.Delete(r.Context(), username)	

	response := WebResponse {
		Code: 200,
		Status: "OK",
	}

	helper.WriteToRequestBody(w, response)
}

func (c *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := params.ByName("username")
	log.Println(username)

	user, err := c.UserService.FindById(r.Context(), username)
	helper.PanicIfError(err)

	response := WebResponse {
		Code: 200,
		Status: "OK",
		Data: user,
	}

	helper.WriteToRequestBody(w, response)

}

// item 

type ItemController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}



type ItemControllerImpl struct {
	ItemService ItemService
}

func NewItemController(itemService ItemService) *ItemControllerImpl {
	return &ItemControllerImpl{
		ItemService: itemService,
	}
}

func (c *ItemControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	itemCreateRequest := ItemCreate{}
	helper.ReadFromRequestBody(r, &itemCreateRequest)

	item := c.ItemService.Create(r.Context(), itemCreateRequest)
	webResponse := WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: item,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c *ItemControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	itemCreateRequest := ItemUpdate{}
	helper.ReadFromRequestBody(r, &itemCreateRequest)

	item := c.ItemService.Update(r.Context(), itemCreateRequest)
	webResponse := WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: item,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c *ItemControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	Id := params.ByName("name")
	
	c.ItemService.Delete(r.Context(), Id)
	webResponse := WebResponse {
		Code: 201,
		Status: "SUCCESS",
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c *ItemControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	Id := params.ByName("name")

	item := c.ItemService.FindById(r.Context(), Id)
	webResponse := WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: item,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c *ItemControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	items := c.ItemService.FindAll(r.Context())
	webResponse := WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: items,
	}

	helper.WriteToRequestBody(w, webResponse)
}

// category

type CategoryHandler interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type CategoryHandlerImpl struct {
	CategoryService		CategoryService
}

func NewCategoryHandler(categoryService CategoryService) CategoryHandler {
	return &CategoryHandlerImpl{
		CategoryService: categoryService,
	}
}


func (h *CategoryHandlerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := CategoryCreate{}
	helper.ReadFromRequestBody(r, &request)

	err := h.CategoryService.Save(r.Context(), request)
	helper.PanicIfError(err)

	response := WebResponse {
		Code: 201,
		Status: "CREATED",
	}

	helper.WriteToRequestBody(w, response)
}

func (h *CategoryHandlerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := CategoryUpdate{}
	helper.ReadFromRequestBody(r, &request)

	id := params.ByName("categoryId")
	request.Id = id

	data, err := h.CategoryService.Update(r.Context(), request)
	helper.PanicIfError(err)

	response := WebResponse {
		Code: 201,
		Status: "OK",
		Data: data,
	}

	helper.WriteToRequestBody(w, response)

}	
