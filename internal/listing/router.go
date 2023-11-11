package listing

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter(user UserController, item ItemController, category CategoryHandler) *httprouter.Router {
	router := httprouter.New()

	// user 
	router.POST("/api/users/signup", user.SignUp)
	router.POST("/api/users/signin", user.SignIn)

	// item
	router.POST("/api/items", item.Create)
	/* 
	router.PUT("/api/items/update/status/:itemId", item.UpdateStatus)
	router.PUT("/api/items/update/description/:itemId", item.UpadteDescription)
	router.PUT("/api/items/update/quantity:itemId", item.UpdateQuantity)
	router.GET("/api/items/:itemId", item.FindById)
	router.GET("/api/items/", item.FindAll)
	router.DELETE("/api/items/:itemId", item.Delete)
	*/

	// category
	router.POST("/api/items/category", category.Create)
	router.PUT("/api/category/:categoryId", category.Update)

	router.PanicHandler = ErrorHandler

	return router
}