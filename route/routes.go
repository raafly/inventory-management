package route

import (
	"github.com/gorilla/mux"
	"github.com/raafly/inventory-management/controller"
)

func NewRouter(user controller.UserControllerImpl, item controller.ItemControllerImpl) *mux.Router{
	r := mux.NewRouter() 

	r.HandleFunc("/api/users", user.SignUp).Methods("POST")
	r.HandleFunc("/api/users", user.SignIn).Methods("GET")
	r.HandleFunc("/api/users/{userId}", user.FindById).Methods("GET").Queries("userId")
	r.HandleFunc("/api/users/", user.FindAll).Methods("GET")
	r.HandleFunc("/apiusers/{userId}", user.Delete).Methods("DELETE").Queries("userId")

	r.HandleFunc("/api/items", item.Create).Methods("POST")
	r.HandleFunc("/api/items", item.Update).Methods("PUT")
	r.HandleFunc("/api/items/{itemId}", item.Delete).Methods("DELETE")
	r.HandleFunc("/api/items/", item.FindAll).Methods("GET")
	r.HandleFunc("/api/items/{itemId}", item.FindById).Methods("GET").Queries("GET")

	return r
}