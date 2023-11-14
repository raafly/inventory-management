package listing

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/raafly/inventory-management/pkg/config"
	"github.com/raafly/inventory-management/pkg/helper"
)

type UserController interface{
	SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImpl struct{
	UserService 	UserService
}

func NewUserController(userService UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c UserControllerImpl) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := UserSignIn{}
	helper.ReadFromRequestBody(r, &userCreateRequest) 

	token, err := c.UserService.SignIn(userCreateRequest)	
	helper.PanicIfError(err)
	webResponse := WebResponse {
		Code: 201,
	}

	var user User
	session, _ := config.Store.Get(r, config.SESSION_ID)

	session.Values["username"] = user.Username
	session.Values["email"] = user.Email

	session.Save(r, w)

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	helper.WriteToRequestBody(w, webResponse)
}

func (c UserControllerImpl) SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := UserSignUp{}
	helper.ReadFromRequestBody(r, &userCreateRequest)

	err := c.UserService.SignUp(userCreateRequest)
	if err != nil {
		fmt.Printf("%v", err)
	}
	webResponse := WebResponse {
		Code: 201,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c UserControllerImpl) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	session, err := config.Store.Get(r, config.SESSION_ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// item 

type ItemController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateStatus(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateQuantity(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpadteDescription(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type ItemControllerImpl struct {
	ItemService ItemService
}

func NewItemController(itemService ItemService) *ItemControllerImpl {
	return &ItemControllerImpl{
		ItemService: itemService,
	}
}

func (c ItemControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	itemCreateRequest := ItemCreate{}
	helper.ReadFromRequestBody(r, &itemCreateRequest)

	c.ItemService.Create(itemCreateRequest)
	webResponse := WebResponse {
		Code: 201,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c ItemControllerImpl) UpdateStatus(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	Id := params.ByName("itemId")
	newId, err := strconv.Atoi(Id) 
	if err != nil {
		panic(err)
	}
	
	itemCreateRequest := ItemUpdate{}
	helper.ReadFromRequestBody(r, &itemCreateRequest)
	
	itemCreateRequest.Id = newId

	c.ItemService.UpdateStatus(itemCreateRequest)
	webResponse := WebResponse {
		Code: 201,
		Data: itemCreateRequest.Id,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c ItemControllerImpl) UpdateQuantity(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	Id := params.ByName("itemId")
	newId, err := strconv.Atoi(Id) 
	if err != nil {
		panic(err)
	}
	
	itemCreateRequest := ItemUpdate{}
	helper.ReadFromRequestBody(r, &itemCreateRequest)
	
	itemCreateRequest.Id = newId
	
	c.ItemService.UpdateQuantity(itemCreateRequest)
	webResponse := WebResponse {
		Code: 200,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c ItemControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	Id := params.ByName("itemId")
	newId, err := strconv.Atoi(Id) 
	if err != nil {
		panic(err)
	}
	
	c.ItemService.Delete(newId)
	webResponse := WebResponse {
		Code: 201,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c ItemControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	Id := params.ByName("itemId")
	newId, err := strconv.Atoi(Id) 
	if err != nil {
		recover()
	}

	item, err := c.ItemService.FindById(newId)
	if err != nil {
		panic(NewNotFoundError(err.Error()))
	}

	webResponse := WebResponse {
		Code: 201,
		Data: item,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c ItemControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	items := c.ItemService.FindAll()
	webResponse := WebResponse {
		Code: 201,
		Data: items,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c ItemControllerImpl) UpadteDescription(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	Id := params.ByName("itemId")
	newId, err := strconv.Atoi(Id) 
	if err != nil {
		panic(err)
	}
	
	itemCreateRequest := ItemUpdate{}
	helper.ReadFromRequestBody(r, &itemCreateRequest)
	
	itemCreateRequest.Id = newId

	c.ItemService.UpadteDescription(itemCreateRequest)
	webResponse := WebResponse {
		Code: 200,
	}

	helper.WriteToRequestBody(w, webResponse)
}

// category

type CategoryHandler interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type CategoryHandlerImpl struct {
	Port CategoryService
}

func NewCategoryHandler(port CategoryService) CategoryHandler {
	return &CategoryHandlerImpl{
		Port: port,
	}
}

func (h CategoryHandlerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := CategoryCreate{}
	helper.ReadFromRequestBody(r, &request)

	h.Port.Save(request)
	response := WebResponse {
		Code: 201,
	}

	helper.WriteToRequestBody(w, response)
}

func (h CategoryHandlerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := CategoryUpdate{}
	helper.ReadFromRequestBody(r, &request)

	id := params.ByName("categoryId")
	request.Id = id

	h.Port.Update(request)

	response := WebResponse {
		Code: 201,
	}

	helper.WriteToRequestBody(w, response)
}	
