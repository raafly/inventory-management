package route

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raafly/inventory-management/controller/port"
)

func NewRouter(user port.UserController, item port.ItemController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users/signup", user.SignUp)
	router.POST("/api/users/signin", user.SignIn)
	router.PUT("/api/users/:username", user.Update)
	router.GET("/api/users/:username", user.FindById)
	router.GET("/api/users/", user.FindAll)
	router.DELETE("/api/users/:username", user.Delete)
	
	router.POST("/api/items", item.Create)
	router.PUT("/api/items/:name", item.Update)
	router.GET("/api/items/:name", item.FindById)
	router.GET("/api/items/", item.FindAll)
	router.DELETE("/api/items/:name", item.Delete)

	return router
}