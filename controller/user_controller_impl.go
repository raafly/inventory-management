package controller

import (
	"net/http"

	"github.com/raafly/inventory-management/helper"
	"github.com/raafly/inventory-management/model"
	"github.com/raafly/inventory-management/service"
)

type UserControllerImpl struct{
	UserService 	service.UserServiceImpl
}

func NewUserController(userService service.UserServiceImpl) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c *UserControllerImpl) SignIn(w http.ResponseWriter, r *http.Request) {
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

func (c *UserControllerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
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

func (c *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	userCreateRequest := model.UserUpdate{}
	helper.ReadFromRequestBody(r, &userCreateRequest)

	c.UserService.Update(r.Context(), userCreateRequest)
	webResponse := model.WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: nil,
	}

	helper.WriteToRequestBody(w, webResponse)	
}

func (c *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	
}

func (c *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {

}

func (c *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	users := c.UserService.FindAll(r.Context())
	webResponse := model.WebResponse {
		Code: 201,
		Status: "SUCCESS",
		Data: users,
	}

	helper.WriteToRequestBody(w, webResponse)
}