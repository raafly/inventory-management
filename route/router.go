package route

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raafly/inventory-management/controller/port"
)

func NewRouter(user port.UserController, item port.ItemController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users", user.SignUp)
	router.GET("/api/users", user.SignIn)
	router.PUT("/api/users/:id", user.Update)
	router.GET("/api/users/:id", user.FindById)
	router.GET("/api/users/", user.FindAll)
	router.DELETE("/api/users/:id", user.Delete)
	
	router.POST("/api/items", item.Create)
	router.PUT("/api/items/:id", item.Update)
	router.GET("/api/items/:id", item.FindById)
	router.GET("/api/items/", item.FindAll)
	router.DELETE("/api/items/:id", item.Delete)

	return router
}