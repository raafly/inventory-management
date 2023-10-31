package listing

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter(user UserController, item ItemController, category CategoryHandler) *httprouter.Router {
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

	// category

	router.POST("/api/items/category", category.Create)

	return router
}