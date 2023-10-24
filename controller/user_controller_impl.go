package controller

import (
	"net/http"
	
	"github.com/julienschmidt/httprouter"
	"github.com/raafly/inventory-management/helper"
	"github.com/raafly/inventory-management/model"
	portService "github.com/raafly/inventory-management/service/port"
)

type UserControllerImpl struct{
	UserService 	portService.UserService
}

func NewUserController(userService portService.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c *UserControllerImpl) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := model.UserSignIn{}
	helper.ReadFromRequestBody(r, &userCreateRequest) 

	user, token := c.UserService.SignIn(r.Context(), userCreateRequest)	
	webResponse := model.WebResponse {
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
	userCreateRequest := model.UserSignUp{}
	helper.ReadFromRequestBody(r, &userCreateRequest)

	user := c.UserService.SignUp(r.Context(), userCreateRequest)
	webResponse := model.WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: user,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (c *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := model.UserUpdate{}
	helper.ReadFromRequestBody(r, &userCreateRequest)

	c.UserService.Update(r.Context(), userCreateRequest)
	webResponse := model.WebResponse {
		Code: 201,
		Status: "SUCCESS",
	}

	helper.WriteToRequestBody(w, webResponse)	
}

func (c *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := params.ByName("username")
	c.UserService.Delete(r.Context(), username)	

	response := model.WebResponse {
		Code: 200,
		Status: "OK",
	}

	helper.WriteToRequestBody(w, response)
}

func (c *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userRequest := model.UserUpdate{}
	helper.ReadFromRequestBody(r, &userRequest)

	username := params.ByName("username")
	c.UserService.FindById(r.Context(), username)

	response := model.WebResponse {
		Code: 200,
		Status: "OK",
	}

	helper.WriteToRequestBody(w, response)

}

func (c *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	users := c.UserService.FindAll(r.Context())
	webResponse := model.WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: users,
	}

	helper.WriteToRequestBody(w, webResponse)
}