package route

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raafly/inventory-management/internal/listing"
)

func NewRouter(user listing.UserController, item listing.ItemController, category listing.CategoryHandler) *httprouter.Router {
	router := httprouter.New()

	// user 
	router.POST("/api/users/signup", user.SignUp)
	router.POST("/api/users/signin", user.SignIn)
	router.POST("/api/users/", user.Logout)

	// item
	router.POST("/api/items", item.Create)
	router.PUT("/api/items/status/:itemId", item.UpdateStatus)
	router.PUT("/api/items/quantity/:itemId", item.UpdateQuantity)
	router.DELETE("/api/items/:itemId", item.Delete)
	router.GET("/api/items/:itemId", item.FindById)
	router.GET("/api/items/", item.FindAll)
	router.PUT("/api/items/description/:itemId", item.UpadteDescription)

	// category
	router.POST("/api/items/category", category.Create)
	router.PUT("/api/category/:categoryId", category.Update)

	router.PanicHandler = listing.ErrorHandler

	return router
}